class TemplateManager:
    """Class to manage PDF templates and themes for character sheets."""

    def __init__(self):
        # Default fantasy theme (for D&D and similar)
        self.fantasy_theme = {
            "primary_color": "#7d5a3c",  # Dark brown
            "secondary_color": "#d3b17d",  # Light brown
            "background_color": "#f0e6d2",  # Beige
            "text_color": "#333",  # Dark gray
            "highlight_color": "#f8f1e3",  # Light beige
            "border_color": "#e6d5b8",  # Medium beige
            "accent_color": "#9c7c54",  # Medium brown
            "header_font": "'Cinzel', serif",
            "body_font": "'Lato', sans-serif",
        }

        # Cyberpunk theme
        self.cyberpunk_theme = {
            "primary_color": "#00ffff",  # Cyan
            "secondary_color": "#ff00ff",  # Magenta
            "background_color": "#0a0a0a",  # Near black
            "text_color": "#ffffff",  # White
            "highlight_color": "#222222",  # Dark gray
            "border_color": "#00ffff",  # Cyan
            "accent_color": "#ff00ff",  # Magenta
            "header_font": "'Orbitron', sans-serif",
            "body_font": "'Roboto Mono', monospace",
        }

        # Steampunk theme
        self.steampunk_theme = {
            "primary_color": "#b87333",  # Copper
            "secondary_color": "#cd7f32",  # Bronze
            "background_color": "#e8d0aa",  # Parchment
            "text_color": "#3a3a3a",  # Dark gray
            "highlight_color": "#f5e6c8",  # Light parchment
            "border_color": "#b87333",  # Copper
            "accent_color": "#5c3317",  # Dark brown
            "header_font": "'IM Fell English', serif",
            "body_font": "'Crimson Text', serif",
        }

        # Gothic horror theme
        self.gothic_theme = {
            "primary_color": "#800020",  # Burgundy
            "secondary_color": "#4a0000",  # Dark red
            "background_color": "#1a1a1a",  # Near black
            "text_color": "#e0e0e0",  # Light gray
            "highlight_color": "#2a2a2a",  # Dark gray
            "border_color": "#800020",  # Burgundy
            "accent_color": "#c0c0c0",  # Silver
            "header_font": "'Playfair Display', serif",
            "body_font": "'Crimson Text', serif",
        }

        # Space opera theme
        self.space_theme = {
            "primary_color": "#4b0082",  # Indigo
            "secondary_color": "#9370db",  # Medium purple
            "background_color": "#0a0a2a",  # Dark blue
            "text_color": "#e0e0e0",  # Light gray
            "highlight_color": "#1a1a4a",  # Slightly lighter blue
            "border_color": "#9370db",  # Medium purple
            "accent_color": "#00bfff",  # Deep sky blue
            "header_font": "'Exo 2', sans-serif",
            "body_font": "'Titillium Web', sans-serif",
        }

        # Theme mapping
        self.theme_map = {
            "cyberpunk": self.cyberpunk_theme,
            "steampunk": self.steampunk_theme,
            "gothic horror": self.gothic_theme,
            "space opera": self.space_theme,
            "fantasy": self.fantasy_theme,
            "default": self.fantasy_theme
        }

        # Font imports for each theme
        self.font_imports = {
            "fantasy": "@import url('https://fonts.googleapis.com/css2?family=Cinzel:wght@400;700&family=Lato:wght@300;400;700&display=swap');",
            "cyberpunk": "@import url('https://fonts.googleapis.com/css2?family=Orbitron:wght@400;700&family=Roboto+Mono:wght@300;400;700&display=swap');",
            "steampunk": "@import url('https://fonts.googleapis.com/css2?family=IM+Fell+English:ital@0;1&family=Crimson+Text:wght@400;600&display=swap');",
            "gothic horror": "@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;700&family=Crimson+Text:wght@400;600&display=swap');",
            "space opera": "@import url('https://fonts.googleapis.com/css2?family=Exo+2:wght@400;700&family=Titillium+Web:wght@300;400;700&display=swap');"
        }

    def __get_theme_css(self, world_theme: str) -> str:
        """Get CSS template based on world theme."""
        theme = self.theme_map.get(world_theme.lower() if world_theme else "default", self.fantasy_theme)
        font_import = self.font_imports.get(world_theme.lower() if world_theme else "default", self.font_imports["fantasy"])

        return f"""
            @page {{
                margin: 0mm;
            }}
            {font_import}

            body {{
                font-family: {theme["body_font"]};
                margin: 0;
                padding: 0;
                background-color: {theme["background_color"]};
                color: {theme["text_color"]};
                line-height: 1.6;
                width: 100%;
                height: 100%;
            }}

            .container {{
                max-width: 80%;
                margin: 20px auto;
                background: {theme["highlight_color"]};
                padding: 30px;
                border-radius: 15px;
                box-shadow: 0 5px 5px rgba(0,0,0,0.1);
                border: 1px solid {theme["secondary_color"]};
            }}

            header {{
                position: relative;
                border-bottom: 2px solid {theme["secondary_color"]};
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
                font-family: {theme["header_font"]};
                color: {theme["primary_color"]};
                margin: 0 0 5px 0;
                font-size: 32px;
                letter-spacing: 1px;
            }}

            .character-subtitle {{
                font-size: 18px;
                color: {theme["accent_color"]};
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
                background: {theme["background_color"]};
                padding: 5px 10px;
                border-radius: 10px;
                border: 1px solid {theme["border_color"]};
                font-size: 10px;
                text-align: center;
            }}

            img {{
                width: 180px;
                height: 180px;
                object-fit: cover;
                border-radius: 10px;
                border: 3px solid {theme["secondary_color"]};
                box-shadow: 0 3px 10px rgba(0,0,0,0.1);
            }}

            .section {{
                margin-bottom: 25px;
                padding: 15px;
                background: {theme["background_color"]};
                border-radius: 10px;
                border-left: 4px solid {theme["secondary_color"]};
            }}

            h2 {{
                font-family: {theme["header_font"]};
                color: {theme["primary_color"]};
                margin-top: 0;
                margin-bottom: 10px;
                padding-bottom: 5px;
                border-bottom: 1px solid {theme["border_color"]};
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
                background: {theme["highlight_color"]};
                padding: 10px 15px;
                border-radius: 8px;
                text-align: center;
                flex: 1;
                min-width: 100px;
                border: 1px solid {theme["secondary_color"]};
            }}

            .stat-label {{
                font-weight: bold;
                font-size: 12px;
                text-transform: uppercase;
                color: {theme["primary_color"]};
                margin-bottom: 5px;
            }}

            .stat-value {{
                font-size: 16px;
            }}

            .highlight {{
                background-color: {theme["highlight_color"]};
                padding: 10px;
                margin: 10px 0;
                border-radius: 8px;
                border-left: 3px solid {theme["secondary_color"]};
            }}

            .footer {{
                text-align: center;
                font-size: 12px;
                color: {theme["accent_color"]};
                margin-top: 20px;
                font-style: italic;
            }}
        """

    def __get_character_image_html(self, character_info: dict) -> str:
        """Get the HTML for the character image."""
        # Get the character image path or use a placeholder
        image_filename = character_info.get('image_filename', 'https://via.placeholder.com/180')

        # Check if the image_filename is a URL or a local file path
        if image_filename.startswith(('http://', 'https://')):
            # If it's a URL, use it directly
            return f'<img src="{image_filename}" alt="Character Portrait"/>'
        else:
            # If it's a local file, use an absolute path
            import os
            absolute_path = os.path.abspath(os.path.join('assets', image_filename))
            return f'<img src="file:///{absolute_path}" alt="Character Portrait"/>'

    def __get_universe_template(self, character_info: dict) -> str:
        """Get HTML template based on world theme."""
        universe = character_info.get('universe', '').lower() if character_info.get('universe') else ''

        # Default D&D template
        if 'd&d' in universe or 'pathfinder' in universe or 'fantasy' in universe:
            return f"""
<html>
<head>
    <style>
        {self.__get_theme_css(character_info.get('world_theme', 'fantasy'))}
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div class="character-title">
                <h1>{character_info.get('name', 'Unknown Name')}</h1>
                <div class="character-subtitle">{character_info.get('race', 'Unknown Race')} • {character_info.get('gender', 'Unknown Gender')}</div>
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
            <p>{character_info['universe_prompt']}</p>

            <div class="stat-block">
                <div class="stat">
                    <div class="stat-label">Setting</div>
                    <div class="stat-value">Fantasy</div>
                </div>
            </div>
        </div>
{'''
        <div class="section">
            <h2>World Theme</h2>
            <p>{character_info['world_theme_prompt']}</p>

            <div class="highlight">
                TODO.
            </div>
        </div>
''' if 'world_theme_prompt' in character_info and character_info['world_theme_prompt'] else ''}

        <div class="footer">
            Character created with LoreCrafter • ID: {character_info.get('id', 'Unknown')}
        </div>
    </div>
</body>
</html>
"""
        # Cyberpunk template (for Cyberpunk RED, Shadowrun, etc.)
        elif 'cyberpunk' in universe or 'shadowrun' in universe:
            return f"""
<html>
<head>
    <style>
        {self.__get_theme_css(character_info.get('world_theme', 'cyberpunk'))}
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div class="character-title">
                <h1>{character_info.get('name', 'Unknown Name')}</h1>
                <div class="character-subtitle">{character_info.get('race', 'Unknown Race')} • {character_info.get('gender', 'Unknown Gender')}</div>
                <div class="character-meta">
                    <div class="meta-item">Universe: {character_info.get('universe', 'Unknown')}</div>
                    <div class="meta-item">Street Cred: High</div>
                </div>
            </div>
            <img src="https://via.placeholder.com/180" alt="Character Portrait"/>
        </header>

        <div class="section">
            <h2>Appearance</h2>
            <p>{character_info['appearance']}</p>

            <div class="stat-block">
                <div class="stat">
                    <div class="stat-label">Cybernetics</div>
                    <div class="stat-value">Unknown</div>
                </div>
            </div>
        </div>

        <div class="section">
            <h2>Personality</h2>
            <p>{character_info['personality']}</p>

            <div class="stat-block">
                <div class="stat">
                    <div class="stat-label">Motivations</div>
                    <div class="stat-value">Survival</div>
                </div>
            </div>
        </div>

        <div class="section">
            <h2>Backstory</h2>
            <p>{character_info['backstory']}</p>

            <div class="highlight">
                TODO
            </div>
        </div>

        <div class="section">
            <h2>Universe</h2>
            <p>{character_info['universe']}</p>

            <div class="stat-block">
                <div class="stat">
                    <div class="stat-label">Setting</div>
                    <div class="stat-value">Dystopian</div>
                </div>
            </div>
        </div>
{'''
        <div class="section">
            <h2>World Theme</h2>
            <p>{character_info['world_theme_prompt']}</p>

            <div class="highlight">
                TODO
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
        # Default template for other universes
        else:
            return f"""
<html>
<head>
    <style>
        {self.__get_theme_css(character_info.get('world_theme', 'default'))}
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div class="character-title">
                <h1>{character_info.get('name', 'Unknown Name')}</h1>
                <div class="character-subtitle">{character_info.get('race', 'Unknown Race')} • {character_info.get('gender', 'Unknown Gender')}</div>
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
                Key moments from {character_info['name']}'s here: TODO.
            </div>
        </div>

        <div class="section">
            <h2>Universe</h2>
            <p>{character_info['universe_prompt']}</p>

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
            <p>{character_info['world_theme_prompt']}</p>

            <div class="highlight">
                TODO
            </div>
        </div>
''' if 'world_theme_prompt' in character_info and character_info['world_theme_prompt'] else ''}

        <div class="footer">
            Character created with LoreCrafter • ID: {character_info.get('id', 'Unknown')}
        </div>
    </div>
</body>
</html>
"""

    def get_template(self, character_info: dict) -> str:
        """Get the appropriate template based on character info."""
        template = self.__get_universe_template(character_info)

        if 'image_filename' in character_info:
            image_html = self.__get_character_image_html(character_info)
            template = template.replace('<img src="https://via.placeholder.com/180" alt="Character Portrait"/>', image_html)

        return template
