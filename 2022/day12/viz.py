from PIL import Image, ImageDraw


last = Image.new('RGB', (41, 93), (0,0,0))
images = [last]

def colorfor(val):
    c = 256 - round(0.18 * val)
    return (c, c, c)

with open('visits.txt', 'r') as f:
    for line in f:
        parts = line.split(", ")
        i = last.copy()
        i.putpixel((int(parts[0]), int(parts[1])), colorfor(int(parts[2])))
        last = i
        images.append(last)


print("Saving")
images[-1].save('animate_last.gif', optimize=False)


