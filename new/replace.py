import os
from fuzzywuzzy import process

def get_file_paths(directory):
    paths = []
    for root, dirs, files in os.walk(directory):
        for file in files:
            p = os.path.join(root, file)
            paths.append(p)
    return paths    

def replace_match(match_tuple, contents):
    (match, replacement) = match_tuple
    return str.replace(contents, match, replacement)

def name_in_line(line, name_to_find):
    split_lines = line.split()
    (lout, confidence) = process.extractOne(name_to_find, split_lines)
    return lout.strip()

def replace_strings_in_text(text_content, strings):
    for pair in strings:
        text_content = replace_match(pair, text_content)
    return text_content

def auto_replace(text_content, original_project_name):
    pass

def replace_strings(project, strings, original_project_name):
    for current_file_path in get_file_paths(project):
        try:
            with open(current_file_path, 'r+', encoding='utf-8') as f:
                text_content = f.read() # read everything in the file
                text_content = replace_strings_in_text(text_content,strings)
                text_content = auto_replace(text_content, original_project_name)
                f.seek(0) # go to the beginning of the file
                f.write(text_content) # write the new line before
                f.truncate() # cut off any remaining text from the old file
        except UnicodeDecodeError:
            pass # Probably non text data, we can skip it
        except Exception as e:
            print(f'Error in replace.replace_strings with file {current_file_path}:', e, 'occurred.')
