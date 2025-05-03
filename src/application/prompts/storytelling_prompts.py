"""
This module contains common storytelling functions used across different packages.
These functions provide descriptions for universes, world themes, and story tones.
"""


def get_universe(universe: str) -> str:
    """Get the context descriptions for different TTRPG universes."""
    universes = {
        "d&d": "D&D universe, adventurers explore a rich fantasy world filled with dragons, ancient magic, and medieval kingdoms full of mystery.",
        "call of cthulhu": "Call of Cthulhu universe plunges characters into Lovecraftian horror, where cosmic secrets unravel sanity and danger lurks in the unknown.",
        "starfinder": "Starfinder universe, science fiction meets fantasy as players journey through galaxies filled with aliens, starships, and arcane technology.",
        "pathfinder": "Pathfinder universe, a deep high-fantasy universe teeming with diverse cultures, sprawling lore, and epic quests driven by heroism and conflict.",
        "vampire the masquerade": "Vampire The Masquerade immerses players in a modern gothic world of vampires, where political intrigue and personal horror shape every dark corner.",
        "chronicles of darkness": "Chronicles of Darkness universe presents a dark reflection of our modern world, where supernatural entities hide in plain sight and ordinary people confront terrifying mysteries.",
        "warhammer 40k": "Warhammer 40K universe depicts a grim, dystopian future where humanity fights for survival against aliens, demons, and heretics in an endless, brutal war across the galaxy.",
        "shadowrun": "Shadowrun universe blends cyberpunk with fantasy in a near-future world where magic has returned, megacorporations rule, and street mercenaries navigate between advanced technology and mystical powers.",
        "traveller": "Traveller universe offers a hard science fiction setting spanning thousands of worlds, where interstellar traders, explorers, and adventurers carve their own path among the stars.",
        "savage worlds": "Savage Worlds universe provides a versatile framework for fast-paced adventures across genres, from pulp heroes to space explorers, with an emphasis on action and heroic deeds.",
        "fate": "FATE universe emphasizes narrative-driven gameplay across infinite settings, where character aspects and story consequences matter more than rigid rules or specific worlds.",
        "blades in the dark": "Blades in the Dark immerses players in the haunted industrial city of Doskvol, where daring scoundrels build criminal enterprises amid supernatural threats and rival gangs.",
        "numenera": "Numenera universe unfolds in Earth's far future, where explorers discover the remnants of ancient, incomprehensible technologies amid the ruins of eight previous civilizations.",
        "star wars rpg": "Star Wars RPG brings the iconic space fantasy to life, with Jedi, smugglers, and rebels fighting against the Empire across exotic planets in a galaxy far, far away.",
        "cyberpunk red": "Cyberpunk RED universe depicts a dystopian future where cybernetic enhancement, corporate domination, and street violence define life in the dangerous urban sprawls of Night City."
    }

    return universes.get(universe.lower(), universe)


def get_world_theme(world_theme: str) -> str:
    """Get thematic keywords and their descriptions that represent different story genres and settings."""
    world_themes = {
        "cyberpunk": "A cyberpunk world, advanced technology collides with societal decay, where neon-lit cities are ruled by powerful corporations.",
        "solarpunk": "A solarpunk setting envisions a hopeful future shaped by ecological balance, community resilience, and sustainable technology.",
        "steampunk": "A steampunk environment where Victorian aesthetics blends with steam-powered inventions, crafting a world of gears, goggles, and bold exploration.",
        "apocalypse": "An apocalypse theme throws characters into a world ravaged by catastrophe, where survival is uncertain and civilization is a memory.",
        "noir": "A noir setting, the story unfolds in a moody atmosphere of crime, moral ambiguity, and shadowy figures working in the dark.",
        "fantasy": "A fantasy universe is filled with magic, mythical creatures, and grand quests, where heroes rise to challenge fate and ancient forces.",
        "space opera": "Space opera delivers an epic journey across galaxies, rich with interstellar politics, starship battles, and legendary destinies.",
        "gothic horror": "Gothic horror evokes eerie dread and beauty, where crumbling castles, haunting secrets, and supernatural terror abound.",
        "mythology": "A mythology theme draws on ancient tales of gods, epic battles, and heroic quests passed down through generations.",
        "urban fantasy": "Urban fantasy that brings the mystical into the modern day, where magic, monsters, and hidden realms exist just beneath the surface of everyday life.",
    }

    return world_themes.get(world_theme.lower(), world_theme)


def get_story_tone(tone: str) -> str:
    """Get the context descriptions for different story tones."""
    tones = {
        "happy": "A happy tone that brings light and joy to the story, focusing on uplifting experiences, triumphs, and heartfelt connections.",
        "sad": "A sad tone that creates emotional depth, revealing stories of loss, introspection, and personal struggle that touch the heart.",
        "dark": "A dark tone that creates a somber atmosphere where characters face grim realities, moral ambiguity, and difficult choices.",
        "hopeful": "A hopeful tone that inspires perseverance, highlighting resilience, growth, and the belief that things can get better against all odds.",
        "mysterious": "A mysterious tone that weaves intrigue and suspense into the story, where secrets unfold slowly and the unknown keeps readers guessing.",
        "epic": "An epic tone that elevates the story to legendary proportions, with grand challenges, heroic deeds, and destinies that shape entire worlds."
    }

    return tones.get(tone.lower(), tone)
