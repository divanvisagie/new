import os

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

def replace_strings(project, strings):
    for current_file_path in get_file_paths(project):
        try:
            with open(current_file_path, 'r+', encoding='utf-8') as f:
                old = f.read() # read everything in the file
                for pair in strings:
                    old = replace_match(pair, old)
                f.seek(0) # go to the beginning of the file
                f.write(old) # write the new line before
                f.truncate() # cut off any remaining text from the old file
        except UnicodeDecodeError:
            pass # Probably non text data, we can skip it
        except Exception as e:
            print(f'Error in replace.replace_strings with file {current_file_path}:', e, 'occurred.')
