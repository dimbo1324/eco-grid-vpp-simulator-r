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
