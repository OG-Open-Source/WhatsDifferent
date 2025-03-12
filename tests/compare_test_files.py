import sys
import os
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))
from whatsdifferent import WhatsDifferent

wd = WhatsDifferent(format='markup')
changes = wd.compare_files("old_file.txt", "setup.py")
print(wd.format_changes(changes))