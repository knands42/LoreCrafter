'use client';

import React, {useEffect, useState} from 'react';
import {useParams, useRouter} from 'next/navigation';
import Link from 'next/link';
import {FaArrowLeft, FaGlobe, FaUser} from 'react-icons/fa';
import Button from '@/components/Button';
import {campaignApi, characterApi, worldApi} from '@/app/api/apiClient';
import {Campaign, Character, World} from '../../api/types';

export default function CampaignDetail() {
  const params = useParams();
  const router = useRouter();
  const campaignId = params.id as string;

  const [campaign, setCampaign] = useState<Campaign | null>(null);
  const [linkedWorld, setLinkedWorld] = useState<World | null>(null);
  const [linkedCharacters, setLinkedCharacters] = useState<Character[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCampaign = async () => {
      try {
        const data = await campaignApi.getById(campaignId);
        setCampaign(data);

        // Fetch linked world if available
        if (data.linked_world_id) {
          try {
            const worldData = await worldApi.getById(data.linked_world_id);
            setLinkedWorld(worldData);
          } catch (err) {
            console.error('Error fetching linked world:', err);
          }
        }

        // Fetch linked characters if available
        if (data.linked_character_ids && data.linked_character_ids.length > 0) {
          const characters: Character[] = [];
          for (const characterId of data.linked_character_ids) {
            try {
              const characterData = await characterApi.getById(characterId);
              characters.push(characterData);
            } catch (err) {
              console.error(`Error fetching character ${characterId}:`, err);
            }
          }
          setLinkedCharacters(characters);
        }
      } catch (err) {
        console.error('Error fetching campaign:', err);
        setError('Failed to load campaign. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    if (campaignId) {
      fetchCampaign();
    }
  }, [campaignId]);

  if (isLoading) {
    return (
      <div className="flex justify-center items-center min-h-[60vh]">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  if (error || !campaign) {
    return (
      <div className="max-w-3xl mx-auto text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-4">
          {error || 'Campaign not found'}
        </h2>
        <Button onClick={() => router.push('/campaigns')}>
          <FaArrowLeft className="mr-2"/>
          Back to Campaigns
        </Button>
      </div>
    );
  }

  // Format hidden elements for better display
  const formattedHiddenElements = campaign.hidden_elements.split('\n').map((line, index) => (
    <div key={index} className="mb-2">
      {line}
    </div>
  ));

  return (
    <div className="max-w-4xl mx-auto">
      {/* Navigation */}
      <div className="mb-6">
        <Button variant="outline" onClick={() => router.push('/campaigns')}>
          <FaArrowLeft className="mr-2"/>
          Back to Campaigns
        </Button>
      </div>

      <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg">
        {/* Header */}
        <div className="px-4 py-5 sm:px-6">
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
            {campaign.name}
          </h1>
          <p className="mt-1 max-w-2xl text-sm text-gray-500 dark:text-gray-400">
            {campaign.universe} • {campaign.world_theme} • {campaign.tone}
          </p>
        </div>

        {/* Campaign Details */}
        <div className="border-t border-gray-200 dark:border-gray-700">
          <dl>
            {/* Campaign Setting */}
            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:px-6">
              <dt className="text-lg font-medium text-gray-900 dark:text-white mb-4">Campaign Setting</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white whitespace-pre-line">
                {campaign.campaign}
              </dd>
            </div>

            {/* Hidden Elements */}
            <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:px-6">
              <dt className="text-lg font-medium text-gray-900 dark:text-white mb-4">Hidden Elements (DM Only)</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white">
                <div className="border-l-2 border-blue-500 pl-4 space-y-2">
                  {formattedHiddenElements}
                </div>
              </dd>
            </div>

            {/* Linked World */}
            {linkedWorld && (
              <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:px-6">
                <dt className="text-lg font-medium text-gray-900 dark:text-white mb-4">Linked World</dt>
                <dd className="mt-1">
                  <div className="bg-white dark:bg-gray-800 shadow overflow-hidden rounded-lg">
                    <div className="px-4 py-4 sm:px-6 flex justify-between items-center">
                      <div>
                        <h3 className="text-lg font-medium text-gray-900 dark:text-white">
                          {linkedWorld.name}
                        </h3>
                        <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                          {linkedWorld.universe} • {linkedWorld.world_theme} • {linkedWorld.tone}
                        </p>
                      </div>
                      <Link href={`/worlds/${linkedWorld.id}`}>
                        <Button variant="outline" size="sm">
                          <FaGlobe className="mr-2"/>
                          View World
                        </Button>
                      </Link>
                    </div>
                  </div>
                </dd>
              </div>
            )}

            {/* Linked Characters */}
            {linkedCharacters.length > 0 && (
              <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:px-6">
                <dt className="text-lg font-medium text-gray-900 dark:text-white mb-4">Linked Characters</dt>
                <dd className="mt-1">
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    {linkedCharacters.map((character) => (
                      <div
                        key={character.id}
                        className="bg-gray-50 dark:bg-gray-900 shadow overflow-hidden rounded-lg"
                      >
                        <div className="px-4 py-4 sm:px-6 flex justify-between items-center">
                          <div>
                            <h3 className="text-md font-medium text-gray-900 dark:text-white">
                              {character.name}
                            </h3>
                            <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                              {character.race} • {character.gender} • {character.universe}
                            </p>
                          </div>
                          <Link href={`/characters/${character.id}`}>
                            <Button variant="outline" size="sm">
                              <FaUser className="mr-2"/>
                              View Character
                            </Button>
                          </Link>
                        </div>
                      </div>
                    ))}
                  </div>
                </dd>
              </div>
            )}

            {/* Universe, World Theme & Tone */}
            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Universe</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {campaign.universe}
              </dd>
            </div>

            <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">World Theme</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {campaign.world_theme}
              </dd>
            </div>

            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Tone</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {campaign.tone}
              </dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  );
}