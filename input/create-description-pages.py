#!/usr/bin/env python3

import csv
import os

# Function to generate file description page content
def generate_file_description(row):
    author = row[9].replace('\n', '; ')
    return f"""Target file name:
WikiCon 2024 - {row[7]}.webm
----

=={{{{int:filedesc}}}}==
{{{{Information
|description={{{{de|1={row[7]}}}}}
|date={row[8]}
|source=Wikimedia Deutschland, WikiCon 2024 Wiesbaden
|author={author}
|permission={{{{en|Authors and persons shown provide permission retrospectively to Wikimedia Deutschland. Please send enquiries to [mailto:info@wikimedia.de Antonia Hohenstein], [https://www.wikimedia.de/ueber-uns/ansprechpartner_innen/ Contact persons]}}}}
{{{{de|Die Autoren und gezeigten Personen erteilen Wikimedia Deutschland nachträglich ihre Zustimmung. Nachfragen bitte an [mailto:info@wikimedia.de Antonia Hohenstein], [https://www.wikimedia.de/ueber-uns/ansprechpartner_innen/ Ansprechpartner*innen]}}}}
|other versions=
}}}}

=={{{{int:license-header}}}}==
{{{{cc-by-sa-4.0}}}}

[[Category:Videos of WikiCon 2024]]
"""

# Script to read CSV and create files
def process_csv(file_path, output_dir):
    with open(file_path, 'r', encoding='utf-8') as csvfile:
        reader = csv.reader(csvfile)
        
        # Ensure output directory exists
        os.makedirs(output_dir, exist_ok=True)

        # Skip header row
        #next(reader)

        for row in reader:
            # Check the criteria
            if row[14] == 'VÖ ok' and len(row[10]) > 4:
                # Generate content
                content = generate_file_description(row)
                
                # Write to file
                output_file = os.path.join(output_dir, f"{row[10]}.wiki.txt")
                with open(output_file, 'w', encoding='utf-8') as f:
                    f.write(content)

# Example usage
csv_file_path = 'upload.csv'  # Replace with the path to your input CSV
output_directory = 'descriptions'  # Replace with your desired output directory
process_csv(csv_file_path, output_directory)

