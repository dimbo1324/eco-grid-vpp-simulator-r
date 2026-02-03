import pytest
from app.core.steam_table import SteamTable


@pytest.mark.parametrize(
    "temp,expected",
    [
        (-10.0, 0.0061),
        (0.0, 0.0061),
        (10.0, None),
        (100.0, 1.01325),
        (374.0, 221.2),
        (2000.0, 500.0),
    ],
)
def test_get_pressure_basic(temp, expected):
    p = SteamTable.get_pressure(temp)
    if expected is not None:
        assert p == pytest.approx(expected)
    else:
        p0 = SteamTable._DATA[0][1]
        p20 = SteamTable._DATA[1][1]
        ratio = (10.0 - 0.0) / (20.0 - 0.0)
        exp = p0 + (p20 - p0) * ratio
        assert p == pytest.approx(exp)


@pytest.mark.parametrize(
    "temp,expected",
    SteamTable._DATA,
)
def test_get_pressure_exact_points(temp, expected):
    assert SteamTable.get_pressure(temp) == pytest.approx(expected)


@pytest.mark.parametrize(
    "temp,expected",
    [
        (100.5, None),
        (373.0, None),
        (float("inf"), 500.0),
        (float("-inf"), 0.0061),
    ],
)
def test_get_pressure_interpolation_and_extremes(temp, expected):
    p = SteamTable.get_pressure(temp)
    if expected is not None:
        assert p == pytest.approx(expected)
    else:
        ratio = (100.5 - 100.0) / (105.0 - 100.0)
        exp = 1.01325 + (1.208 - 1.01325) * ratio
        assert p == pytest.approx(exp)


from hypothesis import given
from hypothesis.strategies import floats


@given(temp=floats(min_value=-1000, max_value=2000))
def test_get_pressure_always_positive_and_clamped(temp):
    p = SteamTable.get_pressure(temp)
    assert p >= 0.0
    assert p <= SteamTable._DATA[-1][1]
