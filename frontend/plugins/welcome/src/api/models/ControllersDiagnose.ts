/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersDiagnose
 */
export interface ControllersDiagnose {
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    department?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    disease?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    doctor?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDiagnose
     */
    patient?: number;
}

export function ControllersDiagnoseFromJSON(json: any): ControllersDiagnose {
    return ControllersDiagnoseFromJSONTyped(json, false);
}

export function ControllersDiagnoseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersDiagnose {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'department': !exists(json, 'department') ? undefined : json['department'],
        'disease': !exists(json, 'disease') ? undefined : json['disease'],
        'doctor': !exists(json, 'doctor') ? undefined : json['doctor'],
        'patient': !exists(json, 'patient') ? undefined : json['patient'],
    };
}

export function ControllersDiagnoseToJSON(value?: ControllersDiagnose | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'department': value.department,
        'disease': value.disease,
        'doctor': value.doctor,
        'patient': value.patient,
    };
}

