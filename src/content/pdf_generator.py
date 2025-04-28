from weasyprint import HTML
from rich import print

# You could also load this from a file or Jinja2 template engine
def get_template(character_info: dict[str, any]) -> str:
    return f"""
<html>
<head>
    <style>
        body {{
            font-family: 'Georgia', serif;
            margin: 5px;
            background-color: #f5f5f5;
        }}
        .container {{
            background: white;
            padding: 5px;
            border-radius: 10px;
        }}
        h1, h2 {{
            color: #2e4053;
            border-bottom: 1px solid #ccc;
            padding-bottom: 5px;
        }}
        .section {{
            margin-bottom: 10px;
        }}
        .stat {{
            background: #e8eaf6;
            padding: 10px;
            border-radius: 5px;
            text-align: center;
        }}
        img {{
            width: 150px;
            height: 150px;
            object-fit: cover;
            border-radius: 10px;
            float: right;
            margin-left: 20px;
        }}
        ul {{
            padding-left: 20px;
        }}
        .highlight {{
            background-color: #ffe0b2;
            padding: 5px;
            margin: 5px 0;
            border-radius: 5px;
        }}
    </style>
</head>
<body>
    <div class="container">
        <h1>{character_info['name']}</h1>
        <img src="https://via.placeholder.com/150" alt="Character Portrait"/>

        <div class="section">
            <h2>Appearance</h2>
            <p>{character_info['appearance']}</p>
        </div>

        <div class="section">
            <h2>Personality</h2>
            <p>{character_info['personality']}</p>
        </div>

        <div class="section">
            <h2>Backstory</h2>
            <p>{character_info['backstory']}</p>
        </div>

        <div class="section">
            <h2>Universe</h2>
            <p>{character_info['universe']}</p>
        </div>

        <div class="section">
            <h2>World Theme</h2>
            <p>{character_info['world_theme']}</p>
        </div>
    </div>
</body>
</html>
"""


# Create the PDF
def create_pdf(content: dict):
    print("\n[bold yellow]Generating your character sheet...[/bold yellow]")
    html_content = get_template(content)
    HTML(string=html_content).write_pdf("assets/character_profile.pdf")
    print("\n[bold yellow]Character sheet ready[/bold yellow]")
