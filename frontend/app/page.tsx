import React from 'react';
import Link from "next/link";

export default function Home() {
  return (
    <div className="space-y-8">
      {/* hero section */}
      <section className="text-center bg-gradient-to-r from-purple-700 to-indigo-800 rounded-lg text-white shadow-xl px-4 py-10 text-inline">
        <h1 className="text-4xl font-extrabold tracking-tight sm:text-5xl lg:text-6xl">
          Welcome to LoreCrafter
        </h1>
        <p className="mt-6 text-xl max-w-3xl mx-auto">
          Create rich backstories for your tabletop role-playing game characters, worlds, and campaigns using AI.
        </p>
        <div>
          <Link href="/character">

          </Link>
          <Link href="">

          </Link>
        </div>
      </section>

      {/*  */}
      <div>
        <section></section>
        <section></section>
        <section></section>
      </div>

      {/*  */}
      <section></section>
    </div>
  );
}
