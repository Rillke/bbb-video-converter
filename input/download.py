#!/usr/bin/env python
import subprocess
import os

# Define the path to the file and the script to be executed
file_path = 'recordings.txt'
basedir = '/home/bigbluebutton'
baseurl = 'https://bbb.wikipedia.de'
script_path = '/opt/bbb-download/scripts/download_presentation.py'

# Open the file and read it line by line
with open(file_path, 'r') as file:
    os.chdir(basedir)
    for line in file:
        # Strip newline characters from the line
        line_content = line.strip()

        # Split line into id -> desired output name
        lcs = line_content.split('|')
        bbb_id = lcs[0]
        file_name = lcs[1]

        # Execute the script with the line content as a command-line argument
        result = subprocess.run(['python', script_path, baseurl, bbb_id], capture_output=True, text=True)

        # Print the output of the script
        print(f"Output for line '{file_name}':\n{result.stdout}")
        if result.stderr:
            print(f"Error for line '{bbb_id}':\n{result.stderr}")

