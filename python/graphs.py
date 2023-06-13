import matplotlib.pyplot as plt
import pandas as pd
import numpy as np

# Plan: three diagrams, one for nothing, one for thick, one for long
# Y-axis: time
# x-axis: combination n/c pairs
# 2 bars per datapoints: trad, weaver

x_points = [(10000, 1), (50000, 10), (100000, 20)]
endpoints = ["nothing", "thick", "long"]
BASE_PATH = "../scripts/results/"

names = ("n=10k, c=1", "n=50k, c=10", "n=100k, c=20")


def build_file_name(app: str, endpoint: str, x_point: tuple[int, int]) -> str:
    return (
        BASE_PATH
        + app
        + "_"
        + str(x_point[0])
        + "_"
        + str(x_point[1])
        + "_"
        + endpoint
        + ".tsv"
    )


def read_data(service: str) -> pd.DataFrame:
    data = pd.DataFrame()
    for x_point in x_points:
        for endpoint in endpoints:
            file_name = build_file_name(service, endpoint, x_point)
            test_run = pd.read_csv(file_name, sep="\t")
            num_cols = len(test_run.columns)
            test_run.insert(num_cols, "endpoint", endpoint)
            test_run.insert(num_cols + 1, "n", x_point[0])
            test_run.insert(num_cols + 2, "c", x_point[1])
            data = pd.concat([data, test_run])
    return data


weaver_df = read_data("weaver")
trad_df = read_data("trad")


def get_time(df: pd.DataFrame, endpoint: str, n: int, c: int) -> int:
    # Use boolean indexing to filter the DataFrame
    filtered_df = df[(df["endpoint"] == endpoint) & (df["n"] == n) & (df["c"] == c)]
    return filtered_df.ttime.sum() / 1000


def get_data_for_endpoint(df: pd.DataFrame, endpoint: str) -> (int, int, int):
    data = (
        get_time(df, endpoint, 10_000, 1),
        get_time(df, endpoint, 50_000, 10),
        get_time(df, endpoint, 100_000, 20),
    )

    return data


def plot_nothing():
    measurements = {
        "traditional": get_data_for_endpoint(trad_df, "nothing"),
        "service_weaver": get_data_for_endpoint(weaver_df, "nothing"),
    }
    x = np.arange(len(names))  # the label locations
    width = 0.25  # the width of the bars
    multiplier = 0

    fig, ax = plt.subplots(layout="constrained")

    for attribute, measurement in measurements.items():
        offset = width * multiplier
        rects = ax.bar(x + offset, measurement, width, label=attribute)
        ax.bar_label(rects, padding=3)
        multiplier += 1

    ax.set_ylabel("Total time for n requests")
    ax.set_title("Nothing Scenario")
    ax.set_xticks(x + width, names)
    ax.legend(loc="upper left", ncols=3)
    plt.title("Nothing Scenario")
    plt.savefig("../paper/figures/nothing_performance.png")


def plot_wide():
    measurements = {
        "traditional": get_data_for_endpoint(trad_df, "thick"),
        "service_weaver": get_data_for_endpoint(weaver_df, "thick"),
    }
    x = np.arange(len(names))  # the label locations
    width = 0.25  # the width of the bars
    multiplier = 0

    fig, ax = plt.subplots(layout="constrained")

    for attribute, measurement in measurements.items():
        offset = width * multiplier
        rects = ax.bar(x + offset, measurement, width, label=attribute)
        ax.bar_label(rects, padding=3)
        multiplier += 1

    ax.set_ylabel("Total time for n requests")
    ax.set_title("Wide Scenario")
    ax.set_xticks(x + width, names)
    ax.legend(loc="upper left", ncols=3)
    plt.title("Wide Scenario")
    plt.savefig("../paper/figures/wide_performance.png")


def plot_long():
    measurements = {
        "traditional": get_data_for_endpoint(trad_df, "long"),
        "service_weaver": get_data_for_endpoint(weaver_df, "long"),
    }
    x = np.arange(len(names))  # the label locations
    width = 0.25  # the width of the bars
    multiplier = 0

    fig, ax = plt.subplots(layout="constrained")

    for attribute, measurement in measurements.items():
        offset = width * multiplier
        rects = ax.bar(x + offset, measurement, width, label=attribute)
        ax.bar_label(rects, padding=3)
        multiplier += 1

    ax.set_ylabel("Total time for n requests")
    ax.set_title("Long Scenario")
    ax.set_xticks(x + width, names)
    ax.legend(loc="upper left", ncols=3)
    # ax.set_ylim(0, 50)
    plt.title("Long Scenario")
    plt.savefig("../paper/figures/long_performance.png")


if __name__ == "__main__":
    plot_nothing()
    plot_wide()
    plot_long()
