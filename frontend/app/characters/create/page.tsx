'use client';

import React, {useState} from 'react';
import {useRouter} from 'next/navigation';
import {FaArrowLeft, FaArrowRight, FaSave} from 'react-icons/fa';
import Input from '@/components/Input';
import TextArea from '@/components/TextArea';
import Select from '@/components/Select';
import Button from '@/components/Button';
import {characterApi} from '@/app/api/apiClient';
import {CreateCharacterRequest, toneOptions, universeOptions, worldThemeOptions} from '../../api/types';

const steps = [
  {id: 'basics', name: 'Basic Information'},
  {id: 'details', name: 'Character Details'},
  {id: 'world', name: 'World Connection'},
  {id: 'review', name: 'Review & Create'},
];

export default function CreateCharacter() {
  const router = useRouter();
  const [currentStep, setCurrentStep] = useState(0);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [formData, setFormData] = useState<CreateCharacterRequest>({
    name: '',
    gender: '',
    race: '',
    personality: '',
    appearance: '',
    universe: '',
    world_theme: '',
    tone: '',
    custom_story: '',
    linked_world_id: '',
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const {name, value} = e.target;
    setFormData((prev) => ({...prev, [name]: value}));
  };

  const handleSelectChange = (name: string) => (value: string) => {
    setFormData((prev) => ({...prev, [name]: value}));
  };

  const nextStep = () => {
    setCurrentStep((prev) => Math.min(prev + 1, steps.length - 1));
  };

  const prevStep = () => {
    setCurrentStep((prev) => Math.max(prev - 1, 0));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      const character = await characterApi.create(formData);
      router.push(`/characters/${character.id}`);
    } catch (err) {
      console.error('Error creating character:', err);
      setError('Failed to create character. Please try again.');
      setIsLoading(false);
    }
  };

  const renderStepContent = () => {
    switch (currentStep) {
      case 0: // Basic Information
        return (
          <div className="space-y-6">
            <Input
              label="Character Name"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              placeholder="Enter character name"
              required
            />
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Input
                label="Gender"
                name="gender"
                value={formData.gender}
                onChange={handleInputChange}
                placeholder="Enter gender"
                required
              />
              <Input
                label="Race"
                name="race"
                value={formData.race}
                onChange={handleInputChange}
                placeholder="Enter race (e.g., Human, Elf, Dwarf)"
                required
              />
            </div>
            <Select
              label="Universe"
              name="universe"
              options={universeOptions}
              value={formData.universe}
              onChange={handleSelectChange('universe')}
              required
            />
          </div>
        );

      case 1: // Character Details
        return (
          <div className="space-y-6">
            <TextArea
              label="Personality"
              name="personality"
              value={formData.personality || ''}
              onChange={handleInputChange}
              placeholder="Describe your character's personality traits (optional)"
              rows={3}
            />
            <TextArea
              label="Appearance"
              name="appearance"
              value={formData.appearance || ''}
              onChange={handleInputChange}
              placeholder="Describe your character's physical appearance (optional)"
              rows={3}
            />
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
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
        );

      case 2: // World Connection
        return (
          <div className="space-y-6">
            <div className="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-4">
              <div className="flex">
                <div>
                  <p className="text-sm text-yellow-700">
                    Optionally, you can link this character to an existing world or provide a custom backstory.
                  </p>
                </div>
              </div>
            </div>

            <Input
              label="Link to Existing World (ID)"
              name="linked_world_id"
              value={formData.linked_world_id || ''}
              onChange={handleInputChange}
              placeholder="Enter world ID (optional)"
            />

            <TextArea
              label="Custom Backstory"
              name="custom_story"
              value={formData.custom_story || ''}
              onChange={handleInputChange}
              placeholder="Provide a custom backstory for your character (optional)"
              rows={6}
            />
          </div>
        );

      case 3: // Review & Create
        return (
          <div className="space-y-6">
            <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg">
              <div className="px-4 py-5 sm:px-6">
                <h3 className="text-lg leading-6 font-medium text-gray-900 dark:text-white">
                  Character Information
                </h3>
                <p className="mt-1 max-w-2xl text-sm text-gray-500 dark:text-gray-400">
                  Review your character details before creating.
                </p>
              </div>
              <div className="border-t border-gray-200 dark:border-gray-700">
                <dl>
                  <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                    <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Name</dt>
                    <dd
                      className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.name}</dd>
                  </div>
                  <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                    <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Gender</dt>
                    <dd
                      className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.gender}</dd>
                  </div>
                  <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                    <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Race</dt>
                    <dd
                      className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.race}</dd>
                  </div>
                  <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                    <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Universe</dt>
                    <dd
                      className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.universe}</dd>
                  </div>
                  <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                    <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">World Theme</dt>
                    <dd
                      className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.world_theme}</dd>
                  </div>
                  <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                    <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Tone</dt>
                    <dd
                      className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.tone}</dd>
                  </div>
                  {formData.personality && (
                    <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Personality</dt>
                      <dd
                        className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.personality}</dd>
                    </div>
                  )}
                  {formData.appearance && (
                    <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Appearance</dt>
                      <dd
                        className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.appearance}</dd>
                    </div>
                  )}
                  {formData.linked_world_id && (
                    <div className="bg-gray-50 dark:bg-gray-900 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Linked World</dt>
                      <dd
                        className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2">{formData.linked_world_id}</dd>
                    </div>
                  )}
                  {formData.custom_story && (
                    <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                      <dt className="text-sm font-medium text-gray-500 dark:text-gray-400">Custom Backstory</dt>
                      <dd
                        className="mt-1 text-sm text-gray-900 dark:text-white sm:mt-0 sm:col-span-2 whitespace-pre-line">{formData.custom_story}</dd>
                    </div>
                  )}
                </dl>
              </div>
            </div>
          </div>
        );

      default:
        return null;
    }
  };

  return (
    <div className="max-w-3xl mx-auto">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Create a New Character</h1>
        <p className="mt-2 text-gray-600 dark:text-gray-400">
          Fill in the details below to create a character with an AI-generated backstory.
        </p>
      </div>

      {/* Step Indicator */}
      <div className="mb-8">
        <nav aria-label="Progress">
          <ol className="flex items-center">
            {steps.map((step, index) => (
              <li key={step.id}
                  className={`relative ${index < steps.length - 1 ? 'pr-8 sm:pr-20' : ''} ${index > 0 ? 'pl-8 sm:pl-20' : ''}`}>
                {index > 0 && (
                  <div className="absolute inset-0 flex items-center" aria-hidden="true">
                    <div className={`h-0.5 w-full ${index <= currentStep ? 'bg-purple-600' : 'bg-gray-200'}`}></div>
                  </div>
                )}
                <div
                  className={`relative flex h-8 w-8 items-center justify-center rounded-full ${
                    index < currentStep
                      ? 'bg-purple-600'
                      : index === currentStep
                        ? 'border-2 border-purple-600 bg-white'
                        : 'border-2 border-gray-300 bg-white'
                  }`}
                >
                  {index < currentStep ? (
                    <svg className="h-5 w-5 text-white" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                      <path fillRule="evenodd"
                            d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
                            clipRule="evenodd"/>
                    </svg>
                  ) : (
                    <span
                      className={`text-sm font-medium ${
                        index === currentStep ? 'text-purple-600' : 'text-gray-500'
                      }`}
                    >
                        {index + 1}
                      </span>
                  )}
                </div>
                <div
                  className="hidden sm:block absolute top-10 left-0 w-32 -ml-16 text-center text-xs font-medium text-gray-500">
                  {step.name}
                </div>
              </li>
            ))}
          </ol>
        </nav>
      </div>

      {/* Form */}
      <form onSubmit={handleSubmit}>
        {error && (
          <div className="mb-4 bg-red-50 border-l-4 border-red-400 p-4">
            <div className="flex">
              <div>
                <p className="text-sm text-red-700">{error}</p>
              </div>
            </div>
          </div>
        )}

        {renderStepContent()}

        <div className="mt-8 flex justify-between">
          <Button
            type="button"
            variant="outline"
            onClick={prevStep}
            disabled={currentStep === 0}
          >
            <FaArrowLeft className="mr-2"/>
            Previous
          </Button>

          {currentStep < steps.length - 1 ? (
            <Button
              type="button"
              onClick={nextStep}
            >
              Next
              <FaArrowRight className="ml-2"/>
            </Button>
          ) : (
            <Button
              type="submit"
              isLoading={isLoading}
            >
              <FaSave className="mr-2"/>
              Create Character
            </Button>
          )}
        </div>
      </form>
    </div>
  );
}