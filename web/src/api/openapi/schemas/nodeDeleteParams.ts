/**
 * Generated by orval v6.28.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { MoveChildNodesQueryParameter } from "./moveChildNodesQueryParameter";
import type { TargetNodeSlugQueryParameter } from "./targetNodeSlugQueryParameter";

export type NodeDeleteParams = {
  /**
   * Where to move child nodes.
   */
  target_node?: TargetNodeSlugQueryParameter;
  /**
   * Whether or not to move child nodes.
   */
  move_child_nodes?: MoveChildNodesQueryParameter;
};