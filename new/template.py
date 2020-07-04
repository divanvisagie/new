import yaml

def get_input(prompt):
    return input(prompt)

def enter_loop_with(config, prompt=get_input):
    replacement_map = []
    for item in config['replace']['strings']:
        match = item['match']
        description = item['description']
        response = prompt(f'Enter replacement text for:\n  {match}\nDescription:\n  {description}\n > ')
        replacement_map.append((match, response))
    return replacement_map

def read(project):
    # Prompts the user for any template replacement requirements
    print(f'reading template for project: {project}')
    try:
        with open(f'{project}/.new.yml') as file:
            f = yaml.load(file, Loader=yaml.FullLoader)
            mp = enter_loop_with(f)
            return mp
    except:
        print('no template config to process')

    