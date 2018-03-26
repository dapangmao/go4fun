import requests
import json
import os
import logging
from tabulate import tabulate
# from subprocess import call
from collections import Counter


class LeetCode(object):

    def __init__(self):
        self.solution = set(os.listdir('./leetcode'))
        self.lang = "go"
        self.travis_ci = "[![Build Status](https://travis-ci.org/dapangmao/go4fun.svg?branch=master)](https://travis-ci.org/dapangmao/go4fun)"

    def get_leetcode(self):
        r = requests.get('https://leetcode.com/api/problems/algorithms')
        j = json.loads(r.text)
        data = j.get('stat_status_pairs')
        all_questions = []
        TRANSLATE = {1: '1-easy', 2: '2-medium', 3: '3-hard'}

        for x in data:
            level = TRANSLATE[x['difficulty']['level']]
            key = x['stat']['question__title']
            qid = x['stat']['question_id']
            slug = x['stat']['question__title_slug']
            total_acs = x['stat']['total_acs']
            total_submitted = x['stat']['total_submitted']
            acceptance = "{0:.0f}%".format(total_acs / total_submitted* 100)
            is_paid = '*' if x['paid_only'] else ''
            url = 'https://leetcode.com/problems/' + slug
            title = '[{}]({})'.format(key, url)
            all_questions.append([key, qid, title, level, is_paid, acceptance])
        return all_questions

    def format_solution_str(self, s):
        url_prefix = 'https://github.com/dapangmao/{}4fun/tree/master/leetcode/'.format(self.lang)
        s = s.replace(' ', '%20')
        return '[solution]({})'.format(url_prefix+s)

    def match_local(self, all_questions):
        all_questions = self.get_leetcode()
        logging.warning("*"*80)
        logging.warning("| There are {} questions, which include {} paid questions.".\
            format(len(all_questions), len([x for x in all_questions if x[4] == '*'])))
        logging.warning("| There are {} solutions.".format(len(self.solution)))
        logging.warning("| " + str(Counter(x[3] for x in all_questions if x[4] != '*')) )
        logging.warning("*"*80)
        for i, x in enumerate(all_questions):
            current = x[0].strip() + '.' + self.lang
            if current in self.solution:
                all_questions[i].append(self.format_solution_str(current))
            else:
                all_questions[i].append('')
        return all_questions

    def write_epub(self, title="leetcode.md"):
        all_questions = self.get_leetcode()
        result = self.match_local(all_questions)
        result.sort(key=lambda x: -x[1])
        out = ""
        for x in result:
            current = x[0] + '.' + self.lang
            if current in self.solution:
                out += "### {}. {}\n\n".format(x[1], x[0])
                out += "- level: {}\n\n".format(x[3])
                out += "```{}\n".format(self.lang)
                with open(os.path.join('leetcode', current), 'r') as infile:
                    for line in infile:
                        out += line
                out += "```\n\n"
        # write to markdown
        with open(title, 'w') as outfile:
            print(out, file=outfile)
        # write to epub
        # call(["pandoc", title, "-o", "leetcode.epub"])

    def write_readme(self):
        all_questions = self.get_leetcode()
        result = self.match_local(all_questions)
        result = [x[1:] for x in result]
        result.sort(key=lambda x: (-x[0], -float(x[4].replace('%', ''))))
        header = ['No', 'Name', 'Difficulty', 'Paid', 'Rate', 'Solution']
        with open('README.md', 'wt') as outfile:
            outfile.write(self.travis_ci + '\n\n')
            outfile.write(tabulate(result, header, tablefmt="pipe"))
