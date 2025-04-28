from weasyprint import HTML
from rich import print

# You could also load this from a file or Jinja2 template engine
def get_template(character_info: dict[str, any]) -> str:
    return f"""
<html>
<head>
    <style>
        @page {{
            margin: 0mm;
        }}
        @import url('https://fonts.googleapis.com/css2?family=Cinzel:wght@400;700&family=Lato:wght@300;400;700&display=swap');

        body {{
            font-family: 'Lato', sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f0e6d2;
            color: #333;
            line-height: 1.6;
            width: 100%;
            height: 100%;
        }}

        .container {{
            max-width: 80%;
            margin: 20px auto;
            background: #fff;
            padding: 30px;
            border-radius: 15px;
            box-shadow: 0 5px 5px rgba(0,0,0,0.1);
            border: 1px solid #d3b17d;
        }}

        header {{
            position: relative;
            border-bottom: 2px solid #d3b17d;
            margin-bottom: 20px;
            padding-bottom: 20px;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }}

        .character-title {{
            flex: 1;
        }}

        h1 {{
            font-family: 'Cinzel', serif;
            color: #7d5a3c;
            margin: 0 0 5px 0;
            font-size: 32px;
            letter-spacing: 1px;
        }}

        .character-subtitle {{
            font-size: 18px;
            color: #9c7c54;
            font-style: italic;
            margin-bottom: 10px;
        }}

        .character-meta {{
            display: flex;
            gap: 15px;
            font-size: 14px;
            align-items: stretch;
        }}

        .meta-item {{
            background: #f8f1e3;
            padding: 5px 10px;
            border-radius: 10px;
            border: 1px solid #e6d5b8;
            font-size: 10px;
            text-align: center;
        }}

        img {{
            width: 180px;
            height: 180px;
            object-fit: cover;
            border-radius: 10px;
            border: 3px solid #d3b17d;
            box-shadow: 0 3px 10px rgba(0,0,0,0.1);
        }}

        .section {{
            margin-bottom: 25px;
            padding: 15px;
            background: #fdfbf7;
            border-radius: 10px;
            border-left: 4px solid #d3b17d;
        }}

        h2 {{
            font-family: 'Cinzel', serif;
            color: #7d5a3c;
            margin-top: 0;
            margin-bottom: 10px;
            padding-bottom: 5px;
            border-bottom: 1px solid #e6d5b8;
        }}

        p {{
            margin: 0;
            text-align: justify;
            font-size: 12px
        }}

        .stat-block {{
            display: flex;
            flex-wrap: wrap;
            gap: 10px;
            margin-top: 15px;
        }}

        .stat {{
            background: #f0e6d2;
            padding: 10px 15px;
            border-radius: 8px;
            text-align: center;
            flex: 1;
            min-width: 100px;
            border: 1px solid #d3b17d;
        }}

        .stat-label {{
            font-weight: bold;
            font-size: 12px;
            text-transform: uppercase;
            color: #7d5a3c;
            margin-bottom: 5px;
        }}

        .stat-value {{
            font-size: 16px;
        }}

        .highlight {{
            background-color: #f8f1e3;
            padding: 10px;
            margin: 10px 0;
            border-radius: 8px;
            border-left: 3px solid #d3b17d;
        }}

        .footer {{
            text-align: center;
            font-size: 12px;
            color: #9c7c54;
            margin-top: 20px;
            font-style: italic;
        }}
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div class="character-title">
                <h1>{character_info.get('name', 'Unknown Name')}</h1>
                <div class="character-subtitle">{character_info.get('race', 'Unknown Race')}</div>
                <div class="character-meta">
                    <div class="meta-item">Universe: {character_info.get('universe', 'Unknown')}</div>
                    <div class="meta-item">Alignment: {character_info.get('alignment', 'Neutral')}</div>
                </div>
            </div>
            <img src="https://via.placeholder.com/180" alt="Character Portrait"/>
        </header>

        <div class="section">
            <h2>Appearance</h2>
            <p>{character_info['appearance']}</p>

            <div class="stat-block">
                <div class="stat">
                    <div class="stat-label">Highlights</div>
                    <div class="stat-value">TODO</div>
                </div>
            </div>
        </div>

        <div class="section">
            <h2>Personality</h2>
            <p>{character_info['personality']}</p>

            <div class="stat-block">
                <div class="stat">
                    <div class="stat-label">Highlights</div>
                    <div class="stat-value">TODO</div>
                </div>
            </div>
        </div>

        <div class="section">
            <h2>Backstory</h2>
            <p>{character_info['backstory']}</p>

            <div class="highlight">
                Key moments from {character_info['name']}'s past have shaped who they are today.
            </div>
        </div>

        <div class="section">
            <h2>Universe</h2>
            <p>{character_info['universe']}</p>

            <div class="stat-block">
                <div class="stat">
                    <div class="stat-label">Setting</div>
                    <div class="stat-value">TODO</div>
                </div>
            </div>
        </div>
{'''
        <div class="section">
            <h2>World Theme</h2>
            <p>{character_info['world_theme']}</p>

            <div class="highlight">

            </div>
        </div>
''' if 'world_theme' in character_info and character_info['world_theme'] else ''}

        <div class="footer">
            Character created with LoreCrafter • ID: {character_info.get('id', 'Unknown')}
        </div>
    </div>
</body>
</html>
"""


# Create the PDF
def create_pdf(content: dict):
    print("\n[bold yellow]Generating your character sheet...[/bold yellow]")
    html_content = get_template(content)
    HTML(string=html_content).write_pdf("assets/character_profile.pdf", presentational_hints=True)
    print("\n[bold yellow]Character sheet ready[/bold yellow]")
