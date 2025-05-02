import type {Metadata} from "next";
import {Geist, Geist_Mono} from "next/font/google";
import "./globals.css";
import Navbar from "@/components/Navbar";
import Footer from "@/components/Footer";
import React from "react";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "LoreCrafter - Create Rich TTRPG Backstories",
  description: "Create detailed characters, immersive worlds, and engaging campaigns for your tabletop role-playing games with AI-powered storytelling.",
  keywords: "TTRPG, RPG, Dungeons and Dragons, D&D, Pathfinder, character creation, world building, campaign creation, AI storytelling",
};

export default function RootLayout(
  {children}: Readonly<{ children: React.ReactNode }>
) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
      <div className="flex flex-col min-h-screen">
          <Navbar />
          <main className="flex-grow bg-gray-50 dark:bg-gray-800">
              <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
                  {children}
              </div>
          </main>
          <Footer />
      </div>
      </body>
    </html>
  );
}
