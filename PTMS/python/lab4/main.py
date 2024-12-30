import pandas as pd
import matplotlib.pyplot as plt
from scipy.stats import pearsonr

inflation_data = pd.read_excel('inflation.xls')
gdp_data = pd.read_excel('gdp.xls')

inflation_data = inflation_data[inflation_data.iloc[:, 0] == "Ukraine"]
gdp_data = gdp_data[gdp_data.iloc[:, 0] == "Ukraine"]

years = range(1980, 2025)
inflation_data = inflation_data.loc[:, [col for col in inflation_data.columns if col in years]]
gdp_data = gdp_data.loc[:, [col for col in gdp_data.columns if col in years]]

inflation_data = inflation_data.T.reset_index()
gdp_data = gdp_data.T.reset_index()

inflation_data.columns = ["Year", "Inflation rate"]
gdp_data.columns = ["Year", "GDP"]

inflation_data["Year"] = pd.to_numeric(inflation_data["Year"], errors="coerce")
inflation_data["Inflation rate"] = pd.to_numeric(inflation_data["Inflation rate"], errors="coerce")

gdp_data["Year"] = pd.to_numeric(gdp_data["Year"], errors="coerce")
gdp_data["GDP"] = pd.to_numeric(gdp_data["GDP"], errors="coerce")

inflation_data = inflation_data.dropna()
gdp_data = gdp_data.dropna()

inflation_data = inflation_data[(inflation_data["Year"] >= 2000) & (inflation_data["Year"] <= 2024)]
gdp_data = gdp_data[(gdp_data["Year"] >= 2000) & (gdp_data["Year"] <= 2024)]

merged_data = pd.merge(inflation_data, gdp_data, on="Year")

corelation, p = pearsonr(merged_data["Inflation rate"], merged_data["GDP"])
print(f"Correlation: {corelation:.2f}")
print(f"P-value: {p:.2f}")

plt.scatter(merged_data["Inflation rate"], merged_data["GDP"], color="green")
plt.title("Inflation rate, average consumer prices/GDP per capita in Ukraine (2000-2024)")
plt.ylabel("U.S. dollars per capita")
plt.xlabel("Annual percent change")
plt.grid()
plt.show()
