import numpy as np
import pandas as pd

# Set the parameters for the dataset
num_points = 300  # Medium size dataset with 300 points
num_dimensions = 2  # Each point has 5 dimensions

# Generate random data points
data = np.random.rand(num_points, num_dimensions)

# Convert the data to a DataFrame
df = pd.DataFrame(data)

# Save the DataFrame to a CSV file
file_path = "./knn_dataset.csv"
df.to_csv(file_path, index=False, header=False)

file_path
