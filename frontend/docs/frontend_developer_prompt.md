# LoreCrafter Frontend Development Prompt

## Project Overview

You are tasked with developing a frontend application for LoreCrafter, a tool that helps users create rich backstories for tabletop role-playing game (TTRPG) characters, worlds, and campaigns using AI. The backend system is already implemented and provides a set of API endpoints that your frontend will interact with.

LoreCrafter allows users to:
1. Create detailed characters with AI-generated backstories, personalities, and appearances
2. Create immersive worlds with histories and timelines
3. Create campaigns that can be linked to worlds and characters
4. Search for existing characters and worlds
5. Generate PDF documents for characters and worlds

## Your Task

Your task is to create a user-friendly, intuitive frontend application that interacts with the LoreCrafter backend API. The frontend should provide a seamless experience for users to create, view, and manage their TTRPG content.

## API Documentation

The backend API is documented in detail in the `api_documentation.md` file. This documentation includes:

- All available endpoints
- Request and response formats for each endpoint
- Data models for characters, worlds, and campaigns
- Error handling information

Please refer to this documentation when implementing the frontend application.

## Frontend Requirements

### General Requirements

1. Create a responsive, user-friendly interface that works well on both desktop and mobile devices
2. Implement proper error handling and user feedback
3. Use the following tech stack:
    - NextJS
    - TailwindCSS
    - TypeScript
    - React Icons
4. Implement proper state management
5. Ensure accessibility compliance

### Specific Features

1. **Character Creation**
   - Create a multi-step form for character creation
   - Allow users to optionally link characters to existing worlds
   - Display the generated character with its backstory, personality, and appearance
   - Show the character image if available
   - Provide an option to download the character as a PDF

2. **World Creation**
   - Create a form for world creation
   - Display the generated world with its history and timeline
   - Show the world image if available
   - Provide an option to download the world as a PDF

3. **Campaign Creation**
   - Create a form for campaign creation
   - Allow users to link the campaign to an existing world
   - Allow users to link existing characters to the campaign
   - Display the generated campaign with its setting and hidden elements

4. **Search Functionality**
   - Implement search for characters and worlds
   - Display search results in a user-friendly way
   - Allow users to view details of search results

5. **Dashboard**
   - Create a dashboard that shows the user's characters, worlds, and campaigns
   - Provide quick access to create new content or view existing content

## Design Considerations

1. **Theme**: The application should have a fantasy/RPG theme
2. **User Experience**: The application should guide users through the creation process with clear instructions and feedback
3. **Visualization**: Consider using creative visualizations for characters, worlds, and campaigns
4. **Responsive Design**: Ensure the application works well on different screen sizes

## Technical Considerations

1. **API Integration**: Use proper API integration techniques (e.g., Axios, Fetch)
2. **State Management**: Use appropriate state management for your chosen framework
3. **Form Handling**: Implement proper form validation and handling
4. **Error Handling**: Implement proper error handling for API requests
5. **Loading States**: Show loading states during API requests

## Deliverables

1. A fully functional frontend application that interacts with the LoreCrafter backend API
2. Source code with clear documentation
3. Setup instructions for local development
4. Any necessary build scripts or configuration files

## Getting Started

1. Review the API documentation in `api_documentation.md`
2. Set up your development environment with your chosen frontend framework
3. Implement the core features one by one, starting with character creation
4. Test your implementation against the backend API
5. Refine the user interface and experience

## Additional Resources

- The backend API is implemented in Python using Typer for CLI commands
- The backend uses OpenAI's API for generating content
- The backend stores characters and worlds in vector databases for semantic search

Good luck with your implementation! If you have any questions about the API or requirements, please refer to the API documentation or ask for clarification.