# importing the libraries
from bs4 import BeautifulSoup
from pandas.core.arrays.string_ import StringArray
import requests
from itertools import islice
from html_table_extractor.extractor import Extractor

url="https://nsscdcl.org/covidbeds/"

# Make a GET request to fetch the raw HTML content
html_content = requests.get(url)

# Parse the html content
soup = BeautifulSoup(html_content.text, "html.parser")
# print(soup.prettify()) # print the parsed data of html

#class="table table-bordered"
div = soup.find_all("div", id = "availventi")
for x in div:
    table = x.find("table",class_="table table-bordered")
#headings = [th.get_text().strip() for th in table.find("tr").find_all("th")]
# print(headings)
datasets = []
for x in table.find_all("td"):
    dataset = x.get_text().strip()
    datasets.append(dataset)
# print(len(datasets))

len_2d = []

for i in range(1,125):
    len_2d.append(2)

def convert(listA, len_2d):
   res = iter(listA)
   return [list(islice(res,i)) for i in len_2d]
res = [convert(datasets, len_2d)]

import csv 
    
# field names 
fields = ['Hospital Name', 'O2 Beds'] 
    
# data rows of csv file 
rows = len_2d
    
# name of csv file 
filename = "./data/Available ventilator.csv"
    
# writing to csv file 
with open(filename, 'w') as csvfile: 
    # creating a csv writer object 
    csvwriter = csv.writer(csvfile) 
        
    # writing the fields 
    csvwriter.writerow(fields) 
        
    for x in res:
        for j in x:
            csvwriter.writerow(j)
