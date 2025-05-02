'use client';

import React, {useState} from 'react';
import {useRouter} from 'next/navigation';
import {FaSave} from 'react-icons/fa';
import Input from '@/components/Input';
import TextArea from '@/components/TextArea';
import Select from '@/components/Select';
import Button from '@/components/Button';
import {worldApi} from '@/app/api/apiClient';
import {CreateWorldRequest, toneOptions, universeOptions, worldThemeOptions} from '../../api/types';

export default function CreateWorld() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [formData, setFormData] = useState<CreateWorldRequest>({
    name: '',
    universe: '',
    world_theme: '',
    tone: '',
    custom_history: '',
    custom_timeline: '',
    custom_backstory: '',
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const {name, value} = e.target;
    setFormData((prev) => ({...prev, [name]: value}));
  };

  const handleSelectChange = (name: string) => (value: string) => {
    setFormData((prev) => ({...prev, [name]: value}));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      const world = await worldApi.create(formData);
      router.push(`/worlds/${world.id}`);
    } catch (err) {
      console.error('Error creating world:', err);
      setError('Failed to create world. Please try again.');
      setIsLoading(false);
    }
  };

  return (
    <div className="max-w-3xl mx-auto">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Create a New World</h1>
        <p className="mt-2 text-gray-600 dark:text-gray-400">
          Fill in the details below to create a world with an AI-generated history and timeline.
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
              label="World Name"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              placeholder="Enter world name"
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
          </div>
        </div>

        <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg p-6">
          <div className="space-y-6">
            <div className="flex items-center justify-between">
              <h2 className="text-xl font-medium text-gray-900 dark:text-white">Custom Details (Optional)</h2>
              <div className="text-sm text-gray-500 dark:text-gray-400">
                All fields below are optional
              </div>
            </div>

            <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-4">
              <div className="flex">
                <div>
                  <p className="text-sm text-yellow-700">
                    You can provide custom details for your world, or leave these fields empty for AI-generated content.
                  </p>
                </div>
              </div>
            </div>

            <TextArea
              label="Custom History"
              name="custom_history"
              value={formData.custom_history || ''}
              onChange={handleInputChange}
              placeholder="Provide a custom history for your world (optional)"
              rows={4}
            />

            <TextArea
              label="Custom Timeline"
              name="custom_timeline"
              value={formData.custom_timeline || ''}
              onChange={handleInputChange}
              placeholder="Provide a custom timeline for your world (optional)"
              rows={4}
            />

            <TextArea
              label="Custom Backstory"
              name="custom_backstory"
              value={formData.custom_backstory || ''}
              onChange={handleInputChange}
              placeholder="Provide additional backstory details for your world (optional)"
              rows={4}
            />
          </div>
        </div>

        <div className="flex justify-end">
          <Button
            type="submit"
            isLoading={isLoading}
          >
            <FaSave className="mr-2"/>
            Create World
          </Button>
        </div>
      </form>
    </div>
  );
}