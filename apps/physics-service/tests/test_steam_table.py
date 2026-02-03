# app/tests/test_steam_table.py
import pytest
from app.core.steam_table import SteamTable


@pytest.mark.parametrize(
    "temp,expected",
    [
        (-10.0, 0.0061),  # below min -> first value
        (0.0, 0.0061),  # exact point
        (10.0, None),  # between 0 and 20 -> interpolate
        (100.0, 1.01325),  # exact point
        (374.0, 221.2),  # exact point near boiling
        (2000.0, 500.0),  # above max -> last value
    ],
)
def test_get_pressure_basic(temp, expected):
    p = SteamTable.get_pressure(temp)
    if expected is not None:
        assert p == pytest.approx(expected)
    else:
        # for 10°C: it's linear between (0,0.0061) and (20,0.0234)
        p0 = SteamTable._DATA[0][1]
        p20 = SteamTable._DATA[1][1]
        ratio = (10.0 - 0.0) / (20.0 - 0.0)
        exp = p0 + (p20 - p0) * ratio
        assert p == pytest.approx(exp)
