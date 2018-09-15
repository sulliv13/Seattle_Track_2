# Data

<p align="center">
  <img src="https://static1.squarespace.com/static/596d24cd4402430bb863ffad/t/5b41e62603ce641f98f2e3cd/1536741696061/?format=1500w" width="350" title="hover text">
</p>


## Example Data Generation Algorithm

<p align="center">
  <img src="https://github.com/FATHOM5/Seattle_Track_2/blob/master/Images/AlgorithmDiagram.png" width="2000" title="hover text">
</p>

### Computationally Efficient Proximity Filtering for AIS Data
1. You can download AIS data from marinecadastre.gov/ais. The data comes in a zipped folder, so you will need to unzip it for use.
2. Depending on how much RAM you have, you may need to split the files into smaller pieces so you can process it. Some of the datasets are 4GB+, and will not run on laptops with only 4GB of RAM. While not necessary for computers with more RAM, users with smaller RAM sizes will need to split the files using a command line utility, or something similar for easier processing.
3. The next step is to split the data into space-time regions, in this case 10 nautical miles in area and 4 hours in height. For this example data a non-overlapping convolution was performed, but an overlapping convolution would provide a more rigorous result and earn you higher points. The data can be split by rounding or truncating the timestamp to the nearest 4 hours, and rounding or truncating the latitude and longitude to the nearest 0.05. Once this code has been computed, the dataset can be sorted and split by it. We can now guarantee that AIS transmissions in the same set are relatively close to one another in both space and time.
4. Naturally, some of these sets will only have a single vessel in them (one unique MMSI). This indicates that only one unique vessel was in the particular 10 square nautical mile block over a period of 4 hours. We can safely delete this part because it means no other vessels were close at all, and there are no possible COLREGS interactions.
5. AIS transmissions happen asynchronously, with vessels sending in data at non-uniform intervals and with non-uniform timing. In order to expedite the computations and simplify the algorithm, we need to figure out where all the vessels are at some uniform sampling rate, in this case 10 minutes. We can accomplish this approximately by averaging a vessels position in 10 minute intervals, giving us the approximate midpoint during that interval.
6. We can then split the data so that AIS transmissions happening in the same now-uniform interval are in the same set.
7. We can now subset all elements in the above set and test to see if they are within 8000 yards within 1 hour, our basic proximity filter settings.

<p align="center">
  <img src="https://static.wixstatic.com/media/3d35e8_2d9eb95a4abe4869afafbf51d29038dc~mv2.png/v1/fill/w_288,h_60,al_c,usm_0.66_1.00_0.01/3d35e8_2d9eb95a4abe4869afafbf51d29038dc~mv2.png" width="150" title="hover text">
</p>
