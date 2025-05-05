# LoreCrafter Improvement Tasks

This document contains a prioritized list of tasks for improving the LoreCrafter project. Each task is marked with a checkbox that can be checked off when completed.

## Core Functionality Improvements

- [x] Create Domains
- [ ] Centralize cli cmds into adapter.in.cli
- [ ] Implement restricted choices for universe, theme, and tone
- [x] Create world lore generation functionality
- [ ] Create a campaign setting (will be only visible for campaign master)
- [ ] Refactor PDF Templates (it's a mess)
- [ ] Add options to create standalone characters tied to a specific world
- [ ] Implement PDF parsing and embedding storage for reference books
- [ ] Enhance prompts with context from relevant embeddings
- [ ] Fix placeholder TODOs in character sheet templates

## Architecture and Code Quality

- [ ] Create a proper project configuration system (replace hardcoded values with config files)
- [ ] Implement proper error handling throughout the application
- [ ] Add logging system for better debugging and monitoring
- [ ] Create unit tests for core functionality
- [ ] Implement integration tests for end-to-end workflows
- [ ] Refactor character generation to use a more modular approach
- [ ] Add type hints consistently across all modules

## User Experience Improvements

- [ ] Enhance PDF character sheets with more detailed information
- [ ] Add support for custom character sheet templates
- [ ] Implement character editing functionality
- [ ] Create a web interface as an alternative to CLI
- [ ] Add support for batch character generation
- [ ] Implement character relationship generation
- [ ] Add progress indicators for long-running operations

## Data Management

- [ ] Implement proper database for character storage (beyond vector store)
- [ ] Add export/import functionality for characters
- [ ] Create backup and restore functionality for the character database
- [ ] Implement versioning for character data
- [ ] Add tagging system for characters

## AI and Content Generation

- [ ] Support multiple LLM providers with fallback options
- [ ] Implement fine-tuning capabilities for specific universes
- [ ] Add more diverse image generation options
- [ ] Create specialized prompt templates for different character types
- [ ] Implement character trait consistency checking
- [ ] Add support for generating character abilities and stats
- [ ] Implement story generation featuring multiple characters

## Documentation and Onboarding

- [ ] Create comprehensive user documentation
- [ ] Add developer documentation with architecture overview
- [ ] Create quickstart guide for new users
- [ ] Add examples of different character types and universes
- [ ] Document API for potential extensions

## Performance and Scalability

- [ ] Optimize LLM prompt usage to reduce token consumption
- [ ] Implement caching for frequently used data
- [ ] Add support for local LLM models
- [ ] Optimize image generation and storage
- [ ] Implement parallel processing for batch operations

## Security and Privacy

- [ ] Implement secure storage for API keys
- [ ] Add user authentication for multi-user setups
- [ ] Create privacy controls for character data
- [ ] Implement data encryption for sensitive information
