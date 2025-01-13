#!/usr/bin/env python3

import csv
import os
import subprocess
from datetime import datetime, timedelta

# Helper function to calculate duration
def calculate_duration(start, end):
    fmt = "%H:%M:%S"
    try:
        start_time = datetime.strptime(start, fmt)
    except ValueError:
        start_time = datetime.strptime("00:00:00", fmt)

    if end:
        try:
            end_time = datetime.strptime(end, fmt)
            duration = end_time - start_time
            return str(duration)
        except ValueError:
            pass

    return None  # No end time provided

# Function to process the CSV and cut videos
def process_videos(csv_file_path, video_file_name_ext):
    with open(csv_file_path, 'r', encoding='utf-8') as csvfile:
        reader = csv.reader(csvfile)

        # Skip header row
        next(reader)

        for row in reader:
            # Check filtering criteria
            if row[14] == 'VÖ ok' and len(row[10]) > 4:
                
                # Get values
                start = row[11] if row[11] else "00:00:00"
                end = row[12] if row[12] else None
                duration = calculate_duration(start, end)

                input_file = f"{row[10]}.{video_file_name_ext}"
                output_file = f"{row[10]}.cut.{video_file_name_ext}"

                # Build ffmpeg command
                if duration:
                    command = [
                        "ffmpeg",
                        "-y",
                        "-ss", start,
                        "-i", input_file,
                        "-t", duration,
                        "-c:a", "copy",
                        "-c:v", "copy",
                        output_file
                    ]
                else:
                    command = [
                        "ffmpeg",
                        "-y",
                        "-ss", start,
                        "-i", input_file,
                        "-c:a", "copy",
                        "-c:v", "copy",
                        output_file
                    ]

                # Run the command
                try:
                    print(f"Executing: {' '.join(command)}")
                    subprocess.run(command, check=True)
                except subprocess.CalledProcessError as e:
                    print(f"Error processing file {input_file}: {e}")

# Example usage
csv_file_path = 'upload.csv'  # Replace with your CSV file path
video_file_name_ext = 'mp4'
process_videos(csv_file_path, video_file_name_ext)

