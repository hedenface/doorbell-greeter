#!/bin/bash

base_dir=$(dirname "$0")

templateFile="$base_dir/template.go"
htmlTemplateFile="$base_dir/../htmlTemplate.go"

cat $templateFile > $htmlTemplateFile

for file in $(find "$base_dir" -name '*.html'); do
    echo "Converting $file"

    # strip everything before the first slash
    # get rid of .html
    # convert - to spaces
    # capitalize first letter of each word
    # get rid of spaces
    name=$(echo "$file" | sed 's|.*/||' | sed 's/\.html//' | sed 's/-/ /g' | sed -e 's/\b./\u\0/g' | sed 's/ //g')
    funcName=$(printf 'html%s' "$name")

    echo "func $funcName() string {" >> $htmlTemplateFile
    echo "return fmt.Sprintf(\"%s\", \`" >> $htmlTemplateFile
    cat "$file" >> $htmlTemplateFile
    echo "\`)" >> $htmlTemplateFile
    echo "}" >> $htmlTemplateFile
    echo "" >> $htmlTemplateFile
done

go fmt $htmlTemplateFile
