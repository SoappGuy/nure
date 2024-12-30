import numpy as np
import matplotlib.pyplot as plt
import time

def uniform_distribution(a, b, size=10000):
    return np.random.uniform(a, b, size)

def exponential_distribution(scale=1.0, size=10000):
    return np.random.exponential(scale, size)

def normal_distribution(mean=0.5, std_dev=0.1, size=10000):
    samples = np.random.normal(mean, std_dev, size)
    return np.clip(samples, 0, 1)

# Normal Distribution (Box-Muller Method)
def normal_distribution_box_muller(mean=0.5, std_dev=0.1, size=10000):
    u1 = np.random.uniform(0, 1, size)
    u2 = np.random.uniform(0, 1, size)
    z0 = np.sqrt(-2 * np.log(u1)) * np.cos(2 * np.pi * u2)
    return mean + std_dev * z0

# Normal Distribution (Ziggurat Algorithm)
def normal_distribution_ziggurat(mean=0.5, std_dev=0.1, size=10000):
    from scipy.stats import norm
    ziggurat = norm.rvs(size=size)
    return mean + std_dev * ziggurat

def plot_histogram(data, bins, color, title, xlabel, ylabel, a, b, cumulative=False):
    plt.hist(data, bins=bins, color=color, cumulative=cumulative, range=(a, b))
    plt.title(title)
    plt.xlabel(xlabel)
    plt.ylabel(ylabel)
    plt.xlim(a, b)

# Function to plot both PDF and CDF
def plot_distribution(pdf_data, cdf_data, a, b, title_pdf, title_cdf, color="blue"):
    plt.figure(figsize=(10, 6))

    # PDF
    plt.subplot(2, 1, 1)
    plot_histogram(
        pdf_data, bins=100, color=color,
        title=title_pdf, xlabel="Value", ylabel="Frequency", a=a, b=b
    )

    # CDF
    plt.subplot(2, 1, 2)
    plot_histogram(
        cdf_data, bins=100, color=color,
        title=title_cdf, xlabel="Value", ylabel="Cumulative Frequency", a=a, b=b, cumulative=True
    )

    plt.tight_layout()
    plt.show()


# Compare performance of Box-Muller and Ziggurat methods
def compare_performance(size=1000000):
    start = time.time()
    _ = normal_distribution_box_muller(size=size)
    box_muller_time = time.time() - start

    start = time.time()
    _ = normal_distribution_ziggurat(size=size)
    ziggurat_time = time.time() - start

    print(f"Box-Muller method time: {box_muller_time:.6f} seconds")
    print(f"Ziggurat algorithm time: {ziggurat_time:.6f} seconds")

if __name__ == "__main__":
    size = 10000

    uniform_data = uniform_distribution(0, 1, size)
    exponential_data = exponential_distribution(1.0, size)
    normal_data = normal_distribution(0.5, 0.1, size)

    plot_distribution(
        uniform_data, uniform_data, 0, 1,
        title_pdf="Uniform Distribution - PDF",
        title_cdf="Uniform Distribution - CDF",
        color="blue"
    )

    plot_distribution(
        exponential_data, exponential_data, 0, np.max(exponential_data),
        title_pdf="Exponential Distribution - PDF",
        title_cdf="Exponential Distribution - CDF",
        color="green"
    )

    plot_distribution(
        normal_data, normal_data, 0, 1,
        title_pdf="Normal Distribution - PDF",
        title_cdf="Normal Distribution - CDF",
        color="red"
    )
