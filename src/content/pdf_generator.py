from weasyprint import HTML
from rich import print

# You could also load this from a file or Jinja2 template engine
html_content = """
<html>
<head>
    <style>
        body {
            font-family: 'Georgia', serif;
            margin: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background: white;
            padding: 20px;
            border-radius: 10px;
        }
        h1, h2 {
            color: #2e4053;
            border-bottom: 1px solid #ccc;
            padding-bottom: 5px;
        }
        .section {
            margin-bottom: 20px;
        }
        .character-sheet {
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 10px;
        }
        .stat {
            background: #e8eaf6;
            padding: 10px;
            border-radius: 5px;
            text-align: center;
        }
        img {
            float: right;
            margin-left: 20px;
            width: 150px;
            height: 150px;
            object-fit: cover;
            border-radius: 10px;
        }
        ul {
            padding-left: 20px;
        }
        .highlight {
            background-color: #ffe0b2;
            padding: 5px;
            margin: 5px 0;
            border-radius: 5px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Thalor Windblade</h1>
        <img src="https://via.placeholder.com/150" alt="Character Portrait"/>

        <div class="section">
            <h2>Appearance</h2>
            <p>Tall and lean elf with silver hair braided with leaves, deep green eyes, and tribal tattoos on his arms.</p>
        </div>

        <div class="section">
            <h2>Personality</h2>
            <ul>
                <li>Brave and selfless</li>
                <li>Silent observer</li>
                <li>Devoted to protecting the forest</li>
            </ul>
        </div>

        <div class="section">
            <h2>Backstory</h2>
            <p>Raised by druids after his village was destroyed, Thalor pledged himself to the ancient spirits of the forest. He now roams the realms as a silent protector, unseen and unstoppable.</p>
        </div>

        <div class="section">
            <h2>World Lore</h2>
            <p>The world of Vhaloria was shattered centuries ago, leaving isolated islands of civilization amidst wild chaos. Ancient magics fuel both wonders and horrors, and alliances are fragile at best.</p>
        </div>

        <div class="section">
            <h2>Highlights</h2>
            <div class="highlight">Defeated the corrupted Ent, Morvath, single-handedly.</div>
            <div class="highlight">Found the lost Amulet of Whispering Leaves.</div>
            <div class="highlight">First elf in 500 years to be named "Guardian of the Verdant Court."</div>
        </div>

        <div class="section">
            <h2>Character Sheet</h2>
            <div class="character-sheet">
                <div class="stat"><strong>Strength:</strong> 14</div>
                <div class="stat"><strong>Dexterity:</strong> 18</div>
                <div class="stat"><strong>Constitution:</strong> 12</div>
                <div class="stat"><strong>Wisdom:</strong> 16</div>
                <div class="stat"><strong>Intelligence:</strong> 13</div>
                <div class="stat"><strong>Charisma:</strong> 11</div>
            </div>
        </div>
    </div>
</body>
</html>
"""

# Create the PDF
def create_pdf(content: dict):
    print("\n[bold yellow]Generating your character sheet...[/bold yellow]")
    HTML(string=html_content).write_pdf("assets/character_profile.pdf")
    print("\n[bold yellow]Character sheet ready[/bold yellow]")
