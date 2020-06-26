#!/usr/bin/python3

import os
import re

kw_url = {}

class Row:
    def __init__(self, id, title, path, keywords=None):
        self.id = id
        self.title = title
        self.path = path
        self.keywords = keywords


def _find_code_file(code_file_path):
    """
    查找指定目录中的代码文件，并迭代返回
    @param    code_file_path 代码文件所在目录
    @return
    """
    fname_rules = r'.*?\.(h|c|hh|cpp|java|py|go)$'
    files = os.listdir(code_file_path)
    for file_name in files:
        if re.match(fname_rules, file_name):
            tmp = os.path.join(code_file_path, file_name)
            yield tmp

def _creat_row(path):
    """
    创建表格一行
    @param    path
    @return
    """
    with open(path, 'r', encoding='utf8') as tmp:
        lines = tmp.readlines()
        t, k = str(lines[5]), str(lines[7])
    p, q = t.find('['), t.find(']')
    id, title = int(t[p+1:q]), t[q+2:-1]
    path = path.replace("\\", "/")
    keywords = k[3:-1].split('|') if len(k) > 2 else None
    return Row(id, title, path, keywords)


def _creat_table(code_file_path):
    """
    创建4栏表格
    @param    code_file_path
    @return
    """
    with open("README.md", 'a', encoding='utf8') as readme:
        readme.write("| ID | Title | Go | Keywords |\n")
        readme.write("| --- | --- | --- | --- |\n")
        row_list = [] 
        for fp in _find_code_file(code_file_path):
            row_list.append(_creat_row(fp))
        row_list = sorted(row_list, key=lambda row: row.id)
        for row in row_list:
            readme.write("| %5d | %s | [Accepted](%s) | " % (row.id, row.title, row.path))
            if row.keywords:
                for kw in row.keywords:
                    if not kw in kw_url:
                        kw_url[kw] = " "
                    readme.write("[%s](%s)、"%(kw, kw_url[kw]))
            readme.write(" |\n")
        readme.write("\n")


if __name__ == '__main__':
    with open("README.md", 'w', encoding='utf8') as readme:
        readme.write("# ultraji的leetcode征程\n\n只为写一份简单易懂的代码。\n\n")
    _creat_table("src")