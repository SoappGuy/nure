import numpy as np
import matplotlib.pyplot as plt
import time

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

def compare_performance(size=1000000):
    start = time.time()
    _ = normal_distribution_box_muller(size=size)
    box_muller_time = time.time() - start

    start = time.time()
    _ = normal_distribution_ziggurat(size=size)
    ziggurat_time = time.time() - start

    return box_muller_time, ziggurat_time

# Visualizing the distributions
def visualize_distributions():
    size = 10000000

    # Generate data
    normal_data_box_muller = normal_distribution_box_muller(0.5, 0.1, size)
    normal_data_ziggurat = normal_distribution_ziggurat(0.5, 0.1, size)

    # Measure performance
    box_muller_time, ziggurat_time = compare_performance(size=size)

    # Plot both distributions on the same graph
    plt.figure(figsize=(10, 6))
    plt.hist(normal_data_box_muller, bins=100, density=True, alpha=0.5, label=f"Box-Muller (Time: {box_muller_time:.6f}s)", color="blue", edgecolor='black')
    plt.hist(normal_data_ziggurat, bins=100, density=True, alpha=0.5, label=f"Ziggurat (Time: {ziggurat_time:.6f}s)", color="green", edgecolor='black')

    plt.title("Normal Distribution Comparison")
    plt.xlabel("Value")
    plt.ylabel("Frequency")
    plt.legend()
    plt.tight_layout()
    plt.show()

# Execute visualization and performance comparison
if __name__ == "__main__":
    visualize_distributions()
    compare_performance()
