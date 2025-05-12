import type {Metadata} from "next";
import {Geist, Geist_Mono} from "next/font/google";
import "./globals.css";
import React from "react";
import {ThemeProvider} from "next-themes";
import {Navbar} from "@/components/Navbar";
import {Footer} from "@/components/Footer";
import {AuthProvider} from "@/context/auth";


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
        // TODO: what suppressHydrationWarning is?
        <html lang="en" suppressHydrationWarning>
        <body
            className={`${geistSans.variable} ${geistMono.variable} antialiased`}
        >
        <ThemeProvider attribute="class" enableSystem defaultTheme="system">
            <AuthProvider>
                <div className="flex flex-col min-h-screen">
                    <Navbar/>
                    <main className="flex-grow bg-gray-50 dark:bg-gray-800">
                        {children}
                    </main>
                    <Footer/>
                </div>
            </AuthProvider>
        </ThemeProvider>
        </body>
        </html>
    );
}
