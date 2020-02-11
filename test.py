import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

# Set up the data pipeline
def get_data(
        location:str = "default/location",
        per: float = 0.2, 
        update: bool = True}
):
    # Function tests the available data locaally
    if os.isfile(os.path.join(pwd, location)):
        file_source = open(os.path.join(pwd, location))
    else:
        raise ValueError("Provided path does not contain files."))

