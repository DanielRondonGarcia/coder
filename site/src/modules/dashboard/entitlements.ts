import type { Entitlements, Feature, FeatureName } from "api/typesGenerated";

/**
 * @param hasLicense true if Enterprise edition
 * @param features record from feature name to feature object
 * @returns record from feature name whether to show the feature
 */
export const getFeatureVisibility = (
	hasLicense: boolean,
	features: Record<string, Feature>,
): Record<string, boolean> => {
	// Always return true for all features to bypass license checks.
	const allFeatures = Object.keys(features).map(feature => [feature, true]);
	return Object.fromEntries(allFeatures);
};

export const selectFeatureVisibility = (
	entitlements: Entitlements,
): Record<FeatureName, boolean> => {
	return getFeatureVisibility(entitlements.has_license, entitlements.features);
};
