from app.core.state import BoilerInputs, BoilerOutputs, BoilerState
import time


def test_dataclasses_default_values():
    inp = BoilerInputs()
    out = BoilerOutputs()
    assert inp.fuel_valve == 0.0
    assert out.furnace_temp == 20.0
    state = BoilerState(timestamp=time.time(), inputs=inp, outputs=out)
    assert isinstance(state.timestamp, float)
    assert state.inputs is inp
    assert state.outputs is out
