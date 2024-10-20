from pprint import pprint

import drawing
import matplotlib.pyplot as plt
import pandas as pd
from numpy import False_
from numpy._core.multiarray import dtype


def main():
    path = "./dataset/tmdb_5000_movies.csv"
    dataset = pd.read_csv(path)

    for series_name, series in dataset.items():
        if "id" not in str(series_name) and series.dtype in [
            dtype("object"),
        ]:
            drawing.draw_pie(series, str(series_name))

    # for series_name, series in dataset.items():
    # if "id" not in str(series_name) and series.dtype in [
    #     dtype("int64"),
    #     dtype("float64"),
    # ]:
    #     drawing.draw_plot(
    #         series.dropna(),
    #         32,
    #         20,
    #         f"Histogram of {series_name}",
    #         "Value",
    #         "Count",
    #     )
    # date = pd.to_datetime(dataset["release_date"])
    # drawing.draw_plot(date)

    # plt.hist(date)
    # drawing.draw_plot(dataset["release_date"])
    # plt.show()


if __name__ == "__main__":
    main()
    input()
