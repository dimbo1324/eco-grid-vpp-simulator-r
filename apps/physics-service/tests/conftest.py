import os
import sys
import pytest

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
if ROOT not in sys.path:
    sys.path.insert(0, ROOT)

from app.core.simulator import BoilerSimulator


@pytest.fixture
def simulator():
    sim = BoilerSimulator()
    sim.last_tick_time = 0.0
    return sim
