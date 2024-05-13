#!/bin/bash

# Directory containing SVG files
SVG_DIR="/Users/yodarango/Desktop/repos/cdn/app/src/ionicons"

# Output JavaScript file
OUTPUT_JS="/Users/yodarango/Desktop/repos/cdn/app/src/JS/icons.js"

# Create or clear the output JavaScript file
> "$OUTPUT_JS"

# Initialize the JavaScript object
echo "const svgMap = {" > "$OUTPUT_JS"

# Iterate over each SVG file in the directory
for svg_file in "$SVG_DIR"/*.svg; do
  # Get the base name of the file (without path and .svg extension)
  base_name=$(basename "$svg_file" .svg)

  # Read the content of the SVG file and escape backticks
  svg_content=$(cat "$svg_file" | sed 's/`/\\`/g' | tr -d '\n')

  # Write the key-value pair to the JavaScript object
  echo "  \"$base_name\": \`$svg_content\`," >> "$OUTPUT_JS"
done

# Close the JavaScript object
echo "};" >> "$OUTPUT_JS"

# Get the size of the output JavaScript file in bytes using macOS specific command
js_file_size_bytes=$(stat -f%z "$OUTPUT_JS")

# Convert the size to megabytes
js_file_size_mb=$(echo "scale=2; $js_file_size_bytes / (1024 * 1024)" | bc)

# Append the total size of the JavaScript file in megabytes at the end
echo "// Total size of this file: ${js_file_size_mb} MB" >> "$OUTPUT_JS"

echo "JavaScript constants have been written to $OUTPUT_JS."