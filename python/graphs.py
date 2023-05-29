# line_plot.py
import matplotlib.pyplot as plt
import pandas as pd

# Plan: two diagrams, one for trad one for weaver
# Y-axis: time
# x-axis: combination n/c pairs
# 3 bars per datapoints: nothing, thick, long

x_points = [(10000, 1), (50000, 10), (100000, 20)]
endpoints = ["nothing", "thick", "long"]
BASE_PATH = "../scripts/results/"


def build_file_name(app: str, endpoint: str, x_point: tuple[int, int]) -> str:
    return (
        BASE_PATH
        + "weaver_"
        + str(x_point[0])
        + "_"
        + str(x_point[1])
        + "_"
        + endpoint
        + ".tsv"
    )


def plot_weaver():
    for x_point in x_points:
        for endpoint in endpoints:
            file_name = build_file_name("weaver", endpoint, x_point)
            data = pd.read_csv(file_name, sep="\t")
            num_cols = len(data.columns)
            data.insert(num_cols - 1, "endpoint", endpoint)
    print(data)


if __name__ == "__main__":
    plot_weaver()
