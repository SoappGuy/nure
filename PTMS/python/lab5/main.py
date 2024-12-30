from numpy import core
import pandas as pd
import matplotlib.pyplot as plt
from scipy.stats import linregress, pearsonr

# Load data
inflation_data = pd.read_excel('inflation.xls')
gdp_data = pd.read_excel('gdp.xls')

# Filter for Ukraine data
inflation_data = inflation_data[inflation_data.iloc[:, 0] == "Ukraine"]
gdp_data = gdp_data[gdp_data.iloc[:, 0] == "Ukraine"]

# Extract data for relevant years
years = range(1980, 2025)
inflation_data = inflation_data.loc[:, [col for col in inflation_data.columns if col in years]]
gdp_data = gdp_data.loc[:, [col for col in gdp_data.columns if col in years]]

# Transpose and reset index
inflation_data = inflation_data.T.reset_index()
gdp_data = gdp_data.T.reset_index()

# Rename columns
inflation_data.columns = ["Year", "Inflation rate"]
gdp_data.columns = ["Year", "GDP"]

# Convert data types
inflation_data["Year"] = pd.to_numeric(inflation_data["Year"], errors="coerce")
inflation_data["Inflation rate"] = pd.to_numeric(inflation_data["Inflation rate"], errors="coerce")

gdp_data["Year"] = pd.to_numeric(gdp_data["Year"], errors="coerce")
gdp_data["GDP"] = pd.to_numeric(gdp_data["GDP"], errors="coerce")

# Drop missing values
inflation_data = inflation_data.dropna()
gdp_data = gdp_data.dropna()

# Filter data for years 2000-2024
inflation_data = inflation_data[(inflation_data["Year"] >= 2000) & (inflation_data["Year"] <= 2024)]
gdp_data = gdp_data[(gdp_data["Year"] >= 2000) & (gdp_data["Year"] <= 2024)]

# Merge datasets
merged_data = pd.merge(inflation_data, gdp_data, on="Year")

# Perform linear regression
x = merged_data["Inflation rate"]
y = merged_data["GDP"]
slope, intercept, r_value, p_value, std_err = linregress(x, y)
y_pred = slope * x + intercept
r_squared = r_value**2

# Display regression equation
print(f"slope:\t\t{slope:.2f}")
print(f"intercept:\t{intercept:.2f}")
print(f"R^2:\t\t{r_squared:.2f}")

print(f"\nGDP = {intercept:.2f} + {slope:.2f} * Inflation rate")

# Plot data and regression line
plt.scatter(merged_data["Inflation rate"], merged_data["GDP"], color="green", label="Data")
plt.plot(merged_data["Inflation rate"], intercept + slope * merged_data["Inflation rate"], color="red", label="Regression line")
plt.xlabel("Inflation rate")
plt.ylabel("GDP")
plt.title("Linear Regression: Inflation rate vs GDP")
plt.legend()
plt.grid()
plt.show()

# Plot residuals
residuals = y - y_pred
plt.scatter(x, residuals, color="blue", label="Residuals")
plt.axhline(0, color='red', linestyle='--', label="Zero line")
plt.xlabel("Inflation rate")
plt.ylabel("Residuals")
plt.title("Residuals of Linear Regression")
plt.legend()
plt.grid()
plt.show()
