/**
 * Generated by orval v6.28.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AssetIDs } from "./assetIDs";
import type { ItemDescription } from "./itemDescription";
import type { ItemName } from "./itemName";
import type { PostContent } from "./postContent";
import type { Properties } from "./properties";
import type { Slug } from "./slug";
import type { Url } from "./url";
import type { Visibility } from "./visibility";

export interface ItemInitialProps {
  asset_ids?: AssetIDs;
  content?: PostContent;
  description: ItemDescription;
  name: ItemName;
  properties?: Properties;
  slug: Slug;
  url?: Url;
  visibility?: Visibility;
}
