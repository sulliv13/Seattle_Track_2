# Data

<p align="center">
  <img src="https://static1.squarespace.com/static/596d24cd4402430bb863ffad/t/5b41e62603ce641f98f2e3cd/1536741696061/?format=1500w" width="350" title="hover text">
</p>

## Interested in Getting the Full Datasets?

**Data Science and Seven Seas** is focused on two areas covered in three UTM Zones.  The first is the west coast of the United States coverd by UTM Zone 11.  The second area is the Caribbean which is encompassed by UTM Zones 17 and 18.  This folder contains a utility `get_data.sh` that will download three zip files from https://marinecadastre.gov/ais/ for December 2017 in the three UTM zones. Usage as follows (note that `user$ ` indicates the terminal prompt not a command to enter). 

 1. After cloning the repository onto your own machine...
 2. Enter the data folder.
 `user$ cd Data`
 3. Change the mode of the script so that it is executable.
 `user$ chmod +x get_data.sh`
 4. Run the script.
 `user$ ./get_data.sh`

The script uses `wget` to download the files, but checks to see if the files are already in your `Data` folder before downloading them. 

From these datasets one of the first challenges is creating an efficient algorithm to find two-ship proximity situations that may represent a COLREGS interaction.  An example algorithm below is one way to do this parsing, but could be much more accurate.

## Example Algorithm for Finding Two-Ship Interactions

<p align="center">
  <img src="https://github.com/FATHOM5/Seattle_Track_2/blob/master/Images/AISAlgorithm.png" width="2000" title="hover text">
</p>

### Computationally Efficient Proximity Filtering for AIS Data
1. You can download AIS data from marinecadastre.gov/ais. The data comes in a zipped folder, so you will need to unzip it for use.
2. Depending on how much RAM you have, you may need to split the files into smaller pieces so you can process it. Some of the datasets are 4GB+, and will not run on laptops with only 4GB of RAM. While not necessary for computers with more RAM, users with smaller RAM sizes will need to split the files using a command line utility, or something similar for easier processing.
3. The next step is to split the data into space-time regions, in this case 10 nautical miles in area and 4 hours in height. For this example data a non-overlapping convolution was performed, but an overlapping convolution would provide a more rigorous result and earn you higher points. The data was split by rounding the timestamp to the nearest 4 hours, and rounding latitude and longitude to the nearest 0.05. Once this code has been computed, the dataset was sorted and split by the space-time code. We can now guarantee that AIS transmissions in the same set are relatively close to one another in both space and time.
4. Naturally, some of these sets will only have a single vessel in them (one unique MMSI). This indicates that only one unique vessel was in the particular 10 square nautical mile block over a period of 4 hours. We can safely delete this part because it means no other vessels were close at all, and there are no possible COLREGS interactions.
5. AIS transmissions happen asynchronously, with vessels sending in data at non-uniform intervals and with non-uniform timing. In order to expedite the computations and simplify the algorithm, we need to figure out where all the vessels are at some uniform sampling rate, in this case 10 minutes. We can accomplish this approximately by averaging a vessels position in 10 minute intervals, giving us the approximate midpoint during that interval.
6. We can then split the data so that AIS transmissions happening in the same now-uniform interval are in the same set.
7. We can now subset all elements in the above set and test to see if they are within 8000 yards within 1 hour, our basic proximity filter settings.

### Possible improvements:
1. Splitting the data up by latitude and longitude is an efficienct way to reduce the number of computations, but it can leave out some data if that data is at the edges of the bounding box. One method to insure accuracy would be to have a sliding bounding box, so that via the overlap no such close components are excluded by box boundaries.
2. Stochastic variations could make this algorithm faster.
3. A recursive divide-and-conquer algorithm could also improve the speed of this computation.


<p align="center">
  <img src="https://static.wixstatic.com/media/3d35e8_2d9eb95a4abe4869afafbf51d29038dc~mv2.png/v1/fill/w_288,h_60,al_c,usm_0.66_1.00_0.01/3d35e8_2d9eb95a4abe4869afafbf51d29038dc~mv2.png" width="150" title="hover text">
</p>
