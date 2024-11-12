/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
The Storyden API does not adhere to semantic versioning but instead applies a rolling strategy with deprecations and minimal breaking changes. This has been done mainly for a simpler development process and it may be changed to a more fixed versioning strategy in the future. Ultimately, the primary way Storyden tracks versions is dates, there are no set release tags currently.

 * OpenAPI spec version: rolling
 */
import type { Asset } from "./asset";
import type { LinkDescription } from "./linkDescription";
import type { LinkDomain } from "./linkDomain";
import type { LinkSlug } from "./linkSlug";
import type { LinkTitle } from "./linkTitle";
import type { Url } from "./url";

export interface LinkReferenceProps {
  description?: LinkDescription;
  domain: LinkDomain;
  favicon_image?: Asset;
  primary_image?: Asset;
  slug: LinkSlug;
  title?: LinkTitle;
  url: Url;
}
