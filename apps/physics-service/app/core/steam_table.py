from typing import List, Tuple


class SteamTable:

    _DATA: List[Tuple[float, float]] = [
        (0.0, 0.0),
        (99.0, 0.0),
        (100.0, 1.0),
        (120.0, 2.0),
        (150.0, 4.7),
        (180.0, 10.0),
        (200.0, 15.5),
        (250.0, 39.7),
        (300.0, 85.8),
        (311.0, 100.0),
        (350.0, 165.0),
        (1500.0, 200.0),
    ]

    @staticmethod
    def get_pressure(temp_c: float) -> float:
        data = SteamTable._DATA

        if temp_c <= data[0][0]:
            return data[0][1]

        for i in range(len(data) - 1):
            t1, p1 = data[i]
            t2, p2 = data[i + 1]

            if t1 <= temp_c <= t2:
                ratio = (temp_c - t1) / (t2 - t1)
                pressure = p1 + (p2 - p1) * ratio
                return pressure

        return data[-1][1]
