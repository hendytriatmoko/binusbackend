import sys
import docx2txt

text = docx2txt.process(sys.argv[1])
print(text)