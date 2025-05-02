'use client';

import React, {useEffect, useState} from 'react';
import {useParams, useRouter} from 'next/navigation';
import Image from 'next/image';
import {FaArrowLeft, FaDownload, FaLink} from 'react-icons/fa';
import Button from '@/components/Button';
import {assetApi, characterApi} from '@/app/api/apiClient';
import {Character} from '../../api/types';

export default function CharacterDetail() {
  const params = useParams();
  const router = useRouter();
  const characterId = params.id as string;

  const [character, setCharacter] = useState<Character | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isPdfLoading, setIsPdfLoading] = useState(false);

  useEffect(() => {
    const fetchCharacter = async () => {
      try {
        const data = await characterApi.getById(characterId);
        setCharacter(data);
      } catch (err) {
        console.error('Error fetching character:', err);
        setError('Failed to load character. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    if (characterId) {
      fetchCharacter();
    }
  }, [characterId]);

  const handleDownloadPdf = async () => {
    if (!character) return;

    setIsPdfLoading(true);
    try {
      const pdfBlob = await characterApi.generatePdf(character.id);

      // Create a URL for the blob
      const url = window.URL.createObjectURL(pdfBlob);

      // Create a temporary link and trigger download
      const link = document.createElement('a');
      link.href = url;
      link.download = `${character.name.replace(/\s+/g, '_')}_character.pdf`;
      document.body.appendChild(link);
      link.click();

      // Clean up
      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
    } catch (err) {
      console.error('Error generating PDF:', err);
      alert('Failed to generate PDF. Please try again.');
    } finally {
      setIsPdfLoading(false);
    }
  };

  if (isLoading) {
    return (
      <div className="flex justify-center items-center min-h-[60vh]">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-purple-500"></div>
      </div>
    );
  }

  if (error || !character) {
    return (
      <div className="max-w-3xl mx-auto text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-4">
          {error || 'Character not found'}
        </h2>
        <Button onClick={() => router.push('/characters')}>
          <FaArrowLeft className="mr-2"/>
          Back to Characters
        </Button>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto">
      {/* Navigation */}
      <div className="mb-6">
        <Button variant="outline" onClick={() => router.push('/characters')}>
          <FaArrowLeft className="mr-2"/>
          Back to Characters
        </Button>
      </div>

      <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg">
        {/* Header */}
        <div className="px-4 py-5 sm:px-6 flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
              {character.name}
            </h1>
            <p className="mt-1 max-w-2xl text-sm text-gray-500 dark:text-gray-400">
              {character.race} • {character.gender} • {character.universe}
            </p>
          </div>
          <Button onClick={handleDownloadPdf} isLoading={isPdfLoading}>
            <FaDownload className="mr-2"/>
            Download PDF
          </Button>
        </div>

        {/* Character Image */}
        {character.image_filename && (
          <div className="border-t border-gray-200 dark:border-gray-700">
            <div className="px-4 py-5 sm:px-6 flex justify-center">
              <div className="relative w-64 h-64 rounded-lg overflow-hidden shadow-lg">
                <Image
                  src={assetApi.getCharacterImage(character.image_filename)}
                  alt={character.name}
                  fill
                  className="object-cover"
                />
              </div>
            </div>
          </div>
        )}

        {/* Character Details */}
        <div className="border-t border-gray-200 dark:border-gray-700">
          <dl>
            {/* Backstory */}
            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Backstory</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2 whitespace-pre-line">
                {character.backstory}
              </dd>
            </div>

            {/* Personality */}
            <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Personality</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {character.personality}
              </dd>
            </div>

            {/* Appearance */}
            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Appearance</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {character.appearance}
              </dd>
            </div>

            {/* World Theme & Tone */}
            <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">World Theme & Tone</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {character.world_theme} • {character.tone}
              </dd>
            </div>

            {/* Linked World */}
            {character.linked_world_id && (
              <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Linked World</dt>
                <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => router.push(`/worlds/${character.linked_world_id}`)}
                  >
                    <FaLink className="mr-2"/>
                    View Linked World
                  </Button>
                </dd>
              </div>
            )}
          </dl>
        </div>
      </div>
    </div>
  );
}