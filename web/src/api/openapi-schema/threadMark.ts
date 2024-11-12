/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */

/**
 * A thread's ID and optional slug separated by a dash = it's unique mark.
This allows endpoints to respond to varying forms of a thread's ID.

For example, given a thread with the ID `cc5lnd2s1s4652adtu50` and the
slug `top-10-movies-thread`, Storyden will understand both the forms:
`cc5lnd2s1s4652adtu50-top-10-movies-thread` and `cc5lnd2s1s4652adtu50`
 as the identifier for that thread.

 */
export type ThreadMark = string;
