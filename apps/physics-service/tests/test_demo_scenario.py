# app/tests/test_demo_scenario.py
import pytest
from app.scenarios.demo import get_demo_controls


@pytest.mark.parametrize(
    "elapsed,step_duration,expected_index",
    [
        (0.0, 10.0, 0),
        (9.9, 10.0, 0),
        (10.0, 10.0, 1),
        (19.9, 10.0, 1),
        (20.0, 10.0, 2),
        (30.0, 10.0, 3),  # 30s+ -> last stage
        (100.0, 10.0, 3),
    ],
)
def test_get_demo_controls_stages(elapsed, step_duration, expected_index):
    stage_name, fuel, water, steam = get_demo_controls(elapsed, step_duration)
    # Simple check: returned tuple length and numeric ranges
    assert isinstance(stage_name, str)
    assert 0.0 <= fuel <= 100.0
    assert 0.0 <= water <= 100.0
    assert 0.0 <= steam <= 100.0
    # verify mapping by comparing with expected call
    # reconstruct stage list same as in source to map index -> values
    stages = [
        ("WARM-UP PHASE    ", 50.0, 10.0, 0.0),
        ("FULL POWER MODE  ", 100.0, 20.0, 0.0),
        ("STEAM BLOWDOWN   ", 100.0, 25.0, 80.0),
        ("SHUTDOWN         ", 0.0, 50.0, 100.0),
    ]
    exp = stages[min(int(elapsed // step_duration), len(stages) - 1)]
    assert (stage_name, fuel, water, steam) == exp
