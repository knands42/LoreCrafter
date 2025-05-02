'use client';

import React, {useState} from 'react';
import {useRouter} from 'next/navigation';
import {FaPlus, FaSave, FaSearch, FaTimes} from 'react-icons/fa';
import Input from '@/components/Input';
import TextArea from '@/components/TextArea';
import Select from '@/components/Select';
import Button from '@/components/Button';
import {campaignApi, characterApi, worldApi} from '@/app/api/apiClient';
import {
  Character,
  CreateCampaignRequest,
  toneOptions,
  universeOptions,
  World,
  worldThemeOptions
} from '../../api/types';

export default function CreateCampaign() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Form data
  const [formData, setFormData] = useState<CreateCampaignRequest>({
    name: '',
    universe: '',
    world_theme: '',
    tone: '',
    custom_campaign: '',
    linked_world_id: '',
    linked_character_ids: [],
  });

  // Search states
  const [worldSearchQuery, setWorldSearchQuery] = useState('');
  const [characterSearchQuery, setCharacterSearchQuery] = useState('');
  const [isSearchingWorlds, setIsSearchingWorlds] = useState(false);
  const [isSearchingCharacters, setIsSearchingCharacters] = useState(false);
  const [worldSearchResults, setWorldSearchResults] = useState<World[]>([]);
  const [characterSearchResults, setCharacterSearchResults] = useState<Character[]>([]);
  const [selectedCharacters, setSelectedCharacters] = useState<Character[]>([]);
  const [selectedWorld, setSelectedWorld] = useState<World | null>(null);

  // Handle input changes
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const {name, value} = e.target;
    setFormData((prev) => ({...prev, [name]: value}));
  };

  const handleSelectChange = (name: string) => (value: string) => {
    setFormData((prev) => ({...prev, [name]: value}));
  };

  // Search for worlds
  const handleWorldSearch = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!worldSearchQuery.trim()) return;

    setIsSearchingWorlds(true);
    try {
      const results = await worldApi.search(worldSearchQuery, 5);
      setWorldSearchResults(results);
    } catch (err) {
      console.error('Error searching worlds:', err);
      setError('Failed to search worlds. Please try again.');
    } finally {
      setIsSearchingWorlds(false);
    }
  };

  // Search for characters
  const handleCharacterSearch = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!characterSearchQuery.trim()) return;

    setIsSearchingCharacters(true);
    try {
      const results = await characterApi.search(characterSearchQuery, 5);
      setCharacterSearchResults(results);
    } catch (err) {
      console.error('Error searching characters:', err);
      setError('Failed to search characters. Please try again.');
    } finally {
      setIsSearchingCharacters(false);
    }
  };

  // Select a world
  const handleSelectWorld = (world: World) => {
    setSelectedWorld(world);
    setFormData((prev) => ({
      ...prev,
      linked_world_id: world.id,
      // Optionally update universe, world_theme, and tone to match the selected world
      universe: world.universe,
      world_theme: world.world_theme,
      tone: world.tone,
    }));
    setWorldSearchResults([]);
    setWorldSearchQuery('');
  };

  // Select a character
  const handleSelectCharacter = (character: Character) => {
    if (!selectedCharacters.some(c => c.id === character.id)) {
      const updatedCharacters = [...selectedCharacters, character];
      setSelectedCharacters(updatedCharacters);
      setFormData((prev) => ({
        ...prev,
        linked_character_ids: updatedCharacters.map(c => c.id)
      }));
    }
    setCharacterSearchResults([]);
    setCharacterSearchQuery('');
  };

  // Remove a character
  const handleRemoveCharacter = (characterId: string) => {
    const updatedCharacters = selectedCharacters.filter(c => c.id !== characterId);
    setSelectedCharacters(updatedCharacters);
    setFormData((prev) => ({
      ...prev,
      linked_character_ids: updatedCharacters.map(c => c.id)
    }));
  };

  // Submit form
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      const campaign = await campaignApi.create(formData);
      router.push(`/campaigns/${campaign.id}`);
    } catch (err) {
      console.error('Error creating campaign:', err);
      setError('Failed to create campaign. Please try again.');
      setIsLoading(false);
    }
  };

  return (
    <div className="max-w-3xl mx-auto">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Create a New Campaign</h1>
        <p className="mt-2 text-gray-600 dark:text-gray-400">
          Fill in the details below to create a campaign with an AI-generated setting and hidden elements.
        </p>
      </div>

      {/* Form */}
      <form onSubmit={handleSubmit} className="space-y-6">
        {error && (
          <div className="mb-4 bg-red-50 border-l-4 border-red-400 p-4">
            <div className="flex">
              <div>
                <p className="text-sm text-red-700">{error}</p>
              </div>
            </div>
          </div>
        )}

        <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg p-6">
          <div className="space-y-6">
            <h2 className="text-xl font-medium text-gray-900 dark:text-white">Basic Information</h2>

            <Input
              label="Campaign Name"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              placeholder="Enter campaign name"
              required
            />

            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <Select
                label="Universe"
                name="universe"
                options={universeOptions}
                value={formData.universe}
                onChange={handleSelectChange('universe')}
                required
              />

              <Select
                label="World Theme"
                name="world_theme"
                options={worldThemeOptions}
                value={formData.world_theme}
                onChange={handleSelectChange('world_theme')}
                required
              />

              <Select
                label="Tone"
                name="tone"
                options={toneOptions}
                value={formData.tone}
                onChange={handleSelectChange('tone')}
                required
              />
            </div>

            <TextArea
              label="Custom Campaign Description"
              name="custom_campaign"
              value={formData.custom_campaign || ''}
              onChange={handleInputChange}
              placeholder="Provide a custom description for your campaign (optional)"
              rows={4}
            />
          </div>
        </div>

        <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg p-6">
          <div className="space-y-6">
            <h2 className="text-xl font-medium text-gray-900 dark:text-white">Link to World</h2>

            {selectedWorld ? (
              <div className="bg-gray-50 dark:bg-gray-700 p-4 rounded-md">
                <div className="flex justify-between items-center">
                  <div>
                    <h3 className="text-md font-medium text-gray-900 dark:text-white">{selectedWorld.name}</h3>
                    <p className="text-sm text-gray-500 dark:text-gray-400">
                      {selectedWorld.universe} • {selectedWorld.world_theme} • {selectedWorld.tone}
                    </p>
                  </div>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => {
                      setSelectedWorld(null);
                      setFormData(prev => ({...prev, linked_world_id: ''}));
                    }}
                  >
                    <FaTimes className="mr-1"/>
                    Remove
                  </Button>
                </div>
              </div>
            ) : (
              <div>
                <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-4">
                  <p className="text-sm text-yellow-700">
                    Optionally, you can link this campaign to an existing world.
                  </p>
                </div>

                <div className="flex flex-col sm:flex-row gap-4 mb-4">
                  <div className="flex-grow">
                    <Input
                      label="Search for a World"
                      value={worldSearchQuery}
                      onChange={(e) => setWorldSearchQuery(e.target.value)}
                      placeholder="Enter world name"
                      className="mb-0"
                    />
                  </div>
                  <div className="flex items-end">
                    <Button
                      type="button"
                      onClick={handleWorldSearch}
                      isLoading={isSearchingWorlds}
                    >
                      <FaSearch className="mr-2"/>
                      Search
                    </Button>
                  </div>
                </div>

                {worldSearchResults.length > 0 && (
                  <div className="mt-4 border border-gray-200 dark:border-gray-700 rounded-md overflow-hidden">
                    <ul className="divide-y divide-gray-200 dark:divide-gray-700">
                      {worldSearchResults.map((world) => (
                        <li
                          key={world.id}
                          className="p-4 hover:bg-gray-50 dark:hover:bg-gray-700 cursor-pointer"
                          onClick={() => handleSelectWorld(world)}
                        >
                          <div className="flex justify-between">
                            <div>
                              <h4 className="text-sm font-medium text-gray-900 dark:text-white">{world.name}</h4>
                              <p className="text-xs text-gray-500 dark:text-gray-400">
                                {world.universe} • {world.world_theme} • {world.tone}
                              </p>
                            </div>
                            <Button size="sm" variant="outline">
                              <FaPlus className="mr-1"/>
                              Select
                            </Button>
                          </div>
                        </li>
                      ))}
                    </ul>
                  </div>
                )}
              </div>
            )}
          </div>
        </div>

        <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg p-6">
          <div className="space-y-6">
            <h2 className="text-xl font-medium text-gray-900 dark:text-white">Link Characters</h2>

            <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-4">
              <p className="text-sm text-yellow-700">
                Optionally, you can link existing characters to this campaign.
              </p>
            </div>

            <div className="flex flex-col sm:flex-row gap-4 mb-4">
              <div className="flex-grow">
                <Input
                  label="Search for Characters"
                  value={characterSearchQuery}
                  onChange={(e) => setCharacterSearchQuery(e.target.value)}
                  placeholder="Enter character name"
                  className="mb-0"
                />
              </div>
              <div className="flex items-end">
                <Button
                  type="button"
                  onClick={handleCharacterSearch}
                  isLoading={isSearchingCharacters}
                >
                  <FaSearch className="mr-2"/>
                  Search
                </Button>
              </div>
            </div>

            {/* Selected Characters */}
            {selectedCharacters.length > 0 && (
              <div className="mt-4 mb-6">
                <h3 className="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Selected Characters:</h3>
                <div className="space-y-2">
                  {selectedCharacters.map((character) => (
                    <div
                      key={character.id}
                      className="flex justify-between items-center bg-gray-50 dark:bg-gray-700 p-3 rounded-md"
                    >
                      <div>
                        <h4 className="text-sm font-medium text-gray-900 dark:text-white">{character.name}</h4>
                        <p className="text-xs text-gray-500 dark:text-gray-400">
                          {character.race} • {character.gender} • {character.universe}
                        </p>
                      </div>
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={() => handleRemoveCharacter(character.id)}
                      >
                        <FaTimes className="mr-1"/>
                        Remove
                      </Button>
                    </div>
                  ))}
                </div>
              </div>
            )}

            {/* Character Search Results */}
            {characterSearchResults.length > 0 && (
              <div className="mt-4 border border-gray-200 dark:border-gray-700 rounded-md overflow-hidden">
                <ul className="divide-y divide-gray-200 dark:divide-gray-700">
                  {characterSearchResults.map((character) => (
                    <li
                      key={character.id}
                      className="p-4 hover:bg-gray-50 dark:hover:bg-gray-700 cursor-pointer"
                      onClick={() => handleSelectCharacter(character)}
                    >
                      <div className="flex justify-between">
                        <div>
                          <h4 className="text-sm font-medium text-gray-900 dark:text-white">{character.name}</h4>
                          <p className="text-xs text-gray-500 dark:text-gray-400">
                            {character.race} • {character.gender} • {character.universe}
                          </p>
                        </div>
                        <Button size="sm" variant="outline">
                          <FaPlus className="mr-1"/>
                          Add
                        </Button>
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
        </div>

        <div className="flex justify-end">
          <Button
            type="submit"
            isLoading={isLoading}
          >
            <FaSave className="mr-2"/>
            Create Campaign
          </Button>
        </div>
      </form>
    </div>
  );
}