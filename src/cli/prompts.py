

def get_universe(universe: str) -> str:
    """Get the context descriptions for different TTRPG universes."""
    universes = {
        "d&d": "D&D universe, adventurers explore a rich fantasy world filled with dragons, ancient magic, and medieval kingdoms full of mystery.",
        "call of cthulhu": "Call of Cthulhu universe plunges characters into Lovecraftian horror, where cosmic secrets unravel sanity and danger lurks in the unknown.",
        "starfinder": "Starfinder universe, science fiction meets fantasy as players journey through galaxies filled with aliens, starships, and arcane technology.",
        "pathfinder": "Pathfinder universe, a deep high-fantasy universe teeming with diverse cultures, sprawling lore, and epic quests driven by heroism and conflict.",
        "vampire the masquerade": "Vampire The Masquerade immerses players in a modern gothic world of vampires, where political intrigue and personal horror shape every dark corner."
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


def get_tone_context(tone: str) -> str:
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