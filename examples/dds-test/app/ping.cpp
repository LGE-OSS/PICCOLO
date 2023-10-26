/**
*  Copyright(c) 2022 ZettaScale Technology and others
*
*   This program and the accompanying materials are made available under the
*   terms of the Eclipse Public License v. 2.0 which is available at
*   http://www.eclipse.org/legal/epl-2.0, or the Eclipse Distribution License
*   v. 1.0 which is available at
*   http://www.eclipse.org/org/documents/edl-v10.php.
*
*   SPDX-License-Identifier: EPL-2.0 OR BSD-3-Clause
*/

#include <vector>
#include <chrono>
#include <algorithm>
#include <iomanip>
#include <iostream>
#include <fstream>
#include <thread>
#include <string>
#include <ctime>
#include <time.h>

#include "dds/dds.hpp"
#include "dds/dds.h"
#include "RoundTrip.hpp"

using namespace std;
using namespace org::eclipse::cyclonedds;

static bool D = true;

static RoundTripModule::DataType roundtrip_msg;

//static vector<unsigned long> payloadSize = {8, 64, 512, 1024, 4096, 1048576, 2764800};
static vector<unsigned long> payloadSize = {8};
static unsigned long numSamples = 10;

int main (int argc, char *argv[])
{
    cout << "=== [Publisher] begin " << endl;
 
    dds::domain::DomainParticipant participant(domain::default_id());

    dds::topic::Topic<RoundTripModule::DataType> topic1(participant, "ping");
    dds::topic::Topic<RoundTripModule::DataType> topic2(participant, "pong");

    dds::pub::Publisher publisher(participant);
    dds::pub::DataWriter<RoundTripModule::DataType> writer(publisher, topic1);

    dds::sub::Subscriber subscriber(participant);
    dds::sub::DataReader<RoundTripModule::DataType> reader(subscriber, topic2);

    
    time_t timer = time(NULL);
    struct tm* t = localtime(&timer);
    char filename[100];
    sprintf(filename, "/app/result/%02d%02d-%02d%02d%02d",
            t->tm_mon + 1, t->tm_mday, t->tm_hour, t->tm_min, t->tm_sec);
    ofstream summaryFile(string(filename).append("_summary.csv"));

    cout << "=== [Publisher] Waiting for subscriber. [domain_id] " << participant->domain_id() << endl;
    while (writer.publication_matched_status().current_count() == 0) {
        this_thread::sleep_for(chrono::milliseconds(20));
    }
    while (reader.subscription_matched_status().current_count() == 0) {
        this_thread::sleep_for(chrono::milliseconds(20));
    }
    
    summaryFile << "=== [Publisher] success finding subscriber." << endl << endl;
    cout << "=== [Publisher] success finding subscriber. " << endl;

    this_thread::sleep_for(chrono::milliseconds(20));

    for (unsigned long size : payloadSize) {
        summaryFile << "payload size : " << size << endl;
        roundtrip_msg.payload(vector<uint8_t>(size, 'a'));
        this_thread::sleep_for(chrono::milliseconds(3000));
        long sum_difference = 0L;
        int success = 0;
        int failure = 0;
        int i = 0;
        for (int i = 1; i <= numSamples; i++) {
            struct timespec begin, end;
            dds::sub::LoanedSamples<RoundTripModule::DataType> samples;

            clock_gettime(CLOCK_MONOTONIC, &begin);
            writer.write(roundtrip_msg);

            while (true) {
                samples = reader.take();
                if (samples.length() > 0) {
                    clock_gettime(CLOCK_MONOTONIC, &end);
                    cout << "receive, i : " << i
                            << ", success : " << ++success << endl;
                    break;
                } else {
                    clock_gettime(CLOCK_MONOTONIC, &end);
                    if ((end.tv_sec - begin.tv_sec) > 60) {
                        summaryFile << "don't received, i : " << i
                                << " , sec: " << (end.tv_sec - begin.tv_sec)
                                << " , fail: " << ++failure << endl;
                        break;
                    }
                }
            }

            long diff = (end.tv_sec - begin.tv_sec) * 1000000000L
                    + end.tv_nsec - begin.tv_nsec;
            if (D) {
                cout << diff << endl;
            }
            summaryFile << diff/2 << endl;
            sum_difference += diff/2;

            this_thread::sleep_for(chrono::milliseconds(600));
            if (success >= numSamples) break;
        }
        if (D) cout << endl;
        summaryFile << sum_difference/numSamples << endl << endl;
        this_thread::sleep_for(chrono::milliseconds(3000));
    }

    cout << "=== [Publisher] Done." << endl;
    summaryFile << "=== [Publisher] Done." << endl;

    if (summaryFile.is_open()) {
       summaryFile.close();
    }
    

    while (true) {
        this_thread::sleep_for(chrono::milliseconds(1000));
    }

    return EXIT_SUCCESS;
}
