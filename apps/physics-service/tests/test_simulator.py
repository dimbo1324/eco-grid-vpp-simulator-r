import pytest
from app.core.simulator import BoilerSimulator
from app.core.steam_table import SteamTable
from app.settings import settings


def test_set_controls_clamps(simulator):
    simulator.set_controls(fuel=-50.0, water=150.0, steam=50.0)
    assert simulator.inputs.fuel_valve == 0.0
    assert simulator.inputs.feedwater_valve == 100.0
    assert simulator.inputs.steam_valve == 50.0


def test_tick_temperature_heating_and_cooling(monkeypatch, simulator):
    simulator.outputs.furnace_temp = 20.0
    simulator.inputs.fuel_valve = 100.0

    monkeypatch.setattr("time.time", lambda: 1.0)
    simulator.last_tick_time = 0.0
    simulator.tick()

    target_temp = settings.AMBIENT_TEMP + settings.MAX_FURNACE_TEMP * (
        simulator.inputs.fuel_valve / 100.0
    )
    expected_change = (target_temp - 20.0) * settings.HEATING_RATE * 1.0
    assert simulator.outputs.furnace_temp == pytest.approx(20.0 + expected_change)

    simulator.outputs.furnace_temp = target_temp + 50.0
    monkeypatch.setattr("time.time", lambda: 2.0)
    simulator.last_tick_time = 1.0
    simulator.tick()
    expected_change2 = (
        (target_temp - (target_temp + 50.0)) * settings.COOLING_RATE * 1.0
    )
    assert simulator.outputs.furnace_temp == pytest.approx(
        (target_temp + 50.0) + expected_change2
    )


def test_tick_pressure_and_flow_and_level(monkeypatch, simulator):
    simulator.outputs.furnace_temp = 120.0
    simulator.inputs.fuel_valve = (
        (simulator.outputs.furnace_temp - settings.AMBIENT_TEMP)
        / settings.MAX_FURNACE_TEMP
    ) * 100.0
    simulator.inputs.steam_valve = 50.0
    simulator.inputs.feedwater_valve = 10.0
    simulator.outputs.drum_level = 500.0

    monkeypatch.setattr("time.time", lambda: 2.0)
    simulator.last_tick_time = 0.0

    dt = 2.0
    base_pressure = SteamTable.get_pressure(simulator.outputs.furnace_temp)
    pressure_loss = simulator.inputs.steam_valve * settings.PRESSURE_DROP_RATE * dt
    expected_pressure = max(0.0, base_pressure - pressure_loss)
    expected_flow = (
        expected_pressure / settings.MAX_PRESSURE
    ) * simulator.inputs.steam_valve

    inflow = simulator.inputs.feedwater_valve * settings.FEEDWATER_RATE * dt
    evaporation = (
        (simulator.outputs.furnace_temp / settings.MAX_FURNACE_TEMP)
        * settings.EVAPORATION_RATE
        * dt
    )
    expected_level = simulator.outputs.drum_level + inflow - evaporation
    expected_level = max(0.0, min(settings.MAX_DRUM_LEVEL, expected_level))

    simulator.tick()

    assert simulator.outputs.steam_pressure == pytest.approx(expected_pressure)
    assert simulator.outputs.steam_flow == pytest.approx(expected_flow)
    assert simulator.outputs.drum_level == pytest.approx(expected_level)


def test_drum_level_clamped_to_bounds(monkeypatch, simulator):
    simulator.outputs.drum_level = 1.0
    simulator.outputs.furnace_temp = 1500.0
    simulator.inputs.feedwater_valve = 0.0

    monkeypatch.setattr("time.time", lambda: 1.0)
    simulator.last_tick_time = 0.0

    simulator.tick()
    assert simulator.outputs.drum_level >= 0.0

    simulator.outputs.drum_level = settings.MAX_DRUM_LEVEL - 1.0
    simulator.inputs.feedwater_valve = 100.0
    monkeypatch.setattr("time.time", lambda: 2.0)
    simulator.last_tick_time = 1.0

    simulator.tick()
    assert simulator.outputs.drum_level <= settings.MAX_DRUM_LEVEL
