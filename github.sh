source ~/py3env/bin/activate
git pull
python write_readme.py
python write_epub.py
git add .
git commit -am "more"
git push
