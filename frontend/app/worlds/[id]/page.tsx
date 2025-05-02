'use client';

import React, {useEffect, useState} from 'react';
import {useParams, useRouter} from 'next/navigation';
import Image from 'next/image';
import {FaArrowLeft, FaDownload} from 'react-icons/fa';
import Button from '@/components/Button';
import {assetApi, worldApi} from '@/app/api/apiClient';
import {World} from '../../api/types';

export default function WorldDetail() {
  const params = useParams();
  const router = useRouter();
  const worldId = params.id as string;

  const [world, setWorld] = useState<World | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isPdfLoading, setIsPdfLoading] = useState(false);

  useEffect(() => {
    const fetchWorld = async () => {
      try {
        const data = await worldApi.getById(worldId);
        setWorld(data);
      } catch (err) {
        console.error('Error fetching world:', err);
        setError('Failed to load world. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    if (worldId) {
      fetchWorld();
    }
  }, [worldId]);

  const handleDownloadPdf = async () => {
    if (!world) return;

    setIsPdfLoading(true);
    try {
      const pdfBlob = await worldApi.generatePdf(world.id);

      // Create a URL for the blob
      const url = window.URL.createObjectURL(pdfBlob);

      // Create a temporary link and trigger download
      const link = document.createElement('a');
      link.href = url;
      link.download = `${world.name.replace(/\s+/g, '_')}_world.pdf`;
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
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-indigo-500"></div>
      </div>
    );
  }

  if (error || !world) {
    return (
      <div className="max-w-3xl mx-auto text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-4">
          {error || 'World not found'}
        </h2>
        <Button onClick={() => router.push('/worlds')}>
          <FaArrowLeft className="mr-2"/>
          Back to Worlds
        </Button>
      </div>
    );
  }

  // Format timeline for better display
  const formattedTimeline = world.timeline.split('\n').map((line, index) => (
    <div key={index} className="mb-2">
      {line}
    </div>
  ));

  return (
    <div className="max-w-4xl mx-auto">
      {/* Navigation */}
      <div className="mb-6">
        <Button variant="outline" onClick={() => router.push('/worlds')}>
          <FaArrowLeft className="mr-2"/>
          Back to Worlds
        </Button>
      </div>

      <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg">
        {/* Header */}
        <div className="px-4 py-5 sm:px-6 flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
              {world.name}
            </h1>
            <p className="mt-1 max-w-2xl text-sm text-gray-500 dark:text-gray-400">
              {world.universe} • {world.world_theme} • {world.tone}
            </p>
          </div>
          <Button onClick={handleDownloadPdf} isLoading={isPdfLoading}>
            <FaDownload className="mr-2"/>
            Download PDF
          </Button>
        </div>

        {/* World Image */}
        {world.image_filename && (
          <div className="border-t border-gray-200 dark:border-gray-700">
            <div className="px-4 py-5 sm:px-6 flex justify-center">
              <div className="relative w-full h-64 md:h-80 rounded-lg overflow-hidden shadow-lg">
                <Image
                  src={assetApi.getWorldImage(world.image_filename)}
                  alt={world.name}
                  fill
                  className="object-cover"
                />
              </div>
            </div>
          </div>
        )}

        {/* World Details */}
        <div className="border-t border-gray-200 dark:border-gray-700">
          <dl>
            {/* History */}
            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:px-6">
              <dt className="text-lg font-medium text-gray-900 dark:text-white mb-4">History</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white whitespace-pre-line">
                {world.history}
              </dd>
            </div>

            {/* Timeline */}
            <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:px-6">
              <dt className="text-lg font-medium text-gray-900 dark:text-white mb-4">Timeline</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white">
                <div className="border-l-2 border-indigo-500 pl-4 space-y-2">
                  {formattedTimeline}
                </div>
              </dd>
            </div>

            {/* World Theme & Tone */}
            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">World Theme</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {world.world_theme}
              </dd>
            </div>

            <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Tone</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {world.tone}
              </dd>
            </div>

            <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Universe</dt>
              <dd className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">
                {world.universe}
              </dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  );
}