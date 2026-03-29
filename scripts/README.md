import os
import sys

def get_project_info():
    """Returns a dictionary with project information."""
    project_info = {
        'name': 'devops-scripts',
        'description': 'A collection of DevOps scripts',
        'author': 'Your Name',
        'email': 'your@email.com',
        'version': '1.0.0'
    }
    return project_info

def get_script_info(script_name):
    """Returns a dictionary with script information."""
    script_info = {
        script_name: {
            'description': 'A DevOps script',
            'author': 'Your Name',
            'email': 'your@email.com',
            'version': '1.0.0'
        }
    }
    return script_info

def print_project_info(project_info):
    """Prints project information."""
    print(f"Project Name: {project_info['name']}")
    print(f"Project Description: {project_info['description']}")
    print(f"Author: {project_info['author']}")
    print(f"Email: {project_info['email']}")
    print(f"Version: {project_info['version']}")

def print_script_info(script_info, script_name):
    """Prints script information."""
    print(f"Script Name: {script_name}")
    print(f"Script Description: {script_info[script_name]['description']}")
    print(f"Author: {script_info[script_name]['author']}")
    print(f"Email: {script_info[script_name]['email']}")
    print(f"Version: {script_info[script_name]['version']}")

if __name__ == "__main__":
    project_info = get_project_info()
    print_project_info(project_info)
    
    script_name = 'script1.py'
    script_info = get_script_info(script_name)
    print_script_info(script_info, script_name)