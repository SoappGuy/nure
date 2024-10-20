import json
from typing import Any

import matplotlib.pyplot as plt
import numpy as np
import numpy.typing as npt
from matplotlib.ticker import FuncFormatter, MaxNLocator


def draw_hist(
    data: Any,
    x_bins: int = 32,
    y_bins: int = 20,
    title: str = "Histogram",
    x_label: str = "Value",
    y_label: str = "Count",
):
    fig, ax = plt.subplots()

    data_max = data.max()
    data_min = data.min()

    ax.hist(data, bins=x_bins)
    ax.grid(True)

    ax.set_title(title)
    ax.set_xlabel(x_label)
    ax.set_ylabel(y_label)

    ax.set_xlim(data_min, data_max)
    ax.set_xticks(np.linspace(data_min, data_max, num=x_bins))
    fig.autofmt_xdate(rotation=60, ha="center")

    # ax.set_xticks(ax.get_xticks(), ax.get_xticklabels(), rotation=60, ha="right")

    ax.xaxis.set_major_formatter(FuncFormatter(lambda x, _: f"{int(x):_}"))
    ax.yaxis.set_major_locator(MaxNLocator(nbins=y_bins))

    fig.set_size_inches(16, 9)
    fig.savefig(f"figs/{title}.png", dpi=100)


def draw_pie(data: Any, title: str = "Pie chart"):
    fig, ax = plt.subplots()

    computed = {}
    for row in data:
        try:
            data = json.loads(row)
        except:  # noqa: E722
            return
        for genre in data:
            if genre["name"] not in computed:
                computed[genre["name"]] = 1
            else:
                computed[genre["name"]] += 1

    labels, values = list(computed.keys()), list(computed.values())
    _sum = sum(values)
    patches, _texts = ax.pie(values)

    legend_labels = [f"{l} ({(v / _sum * 100):.1f}%)" for l, v in zip(labels, values)]
    ax.legend(
        patches,
        legend_labels,
        loc="center left",
        bbox_to_anchor=(1, 0.5),
    )

    fig.set_size_inches(16, 9)
    fig.savefig(f"figs/{title}.png", dpi=100)
