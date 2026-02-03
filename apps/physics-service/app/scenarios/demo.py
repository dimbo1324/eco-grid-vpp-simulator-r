from typing import Tuple


def get_demo_controls(
    elapsed: float, step_duration: float
) -> Tuple[str, float, float, float]:

    step = int(elapsed // step_duration)

    stages = [
        ("РАЗОГРЕВ     ", 50.0, 10.0, 0.0),
        ("ФОРСАЖ       ", 100.0, 20.0, 0.0),
        ("ОТДАЧА ПАРА  ", 100.0, 25.0, 80.0),
        ("ОСТАНОВ      ", 0.0, 50.0, 100.0),
    ]

    return stages[min(step, len(stages) - 1)]
